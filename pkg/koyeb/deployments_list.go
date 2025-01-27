package koyeb

import (
	"strconv"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/errors"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *DeploymentHandler) List(ctx *CLIContext, cmd *cobra.Command, args []string) error {
	list := []koyeb.DeploymentListItem{}

	appId := ""
	if app, _ := cmd.Flags().GetString("app"); app != "" {
		var err error
		if appId, err = h.ResolveAppArgs(ctx, app); err != nil {
			return err
		}
	}
	serviceId := ""
	if service, _ := cmd.Flags().GetString("service"); service != "" {
		var err error
		if serviceId, err = h.ResolveServiceArgs(ctx, service); err != nil {
			return err
		}
	}

	page := int64(0)
	offset := int64(0)
	limit := int64(100)
	for {
		req := ctx.Client.DeploymentsApi.ListDeployments(ctx.Context)

		if appId != "" {
			req = req.AppId(appId)
		}
		if serviceId != "" {
			req = req.ServiceId(serviceId)
		}

		res, resp, err := req.Limit(strconv.FormatInt(limit, 10)).Offset(strconv.FormatInt(offset, 10)).Execute()
		if err != nil {
			return errors.NewCLIErrorFromAPIError(
				"Error while listing the deployments",
				err,
				resp,
			)
		}
		list = append(list, res.GetDeployments()...)

		page++
		offset = page * limit
		if offset >= res.GetCount() {
			break
		}
	}

	full := GetBoolFlags(cmd, "full")
	listDeploymentsReply := NewListDeploymentsReply(ctx.Mapper, &koyeb.ListDeploymentsReply{Deployments: list}, full)
	ctx.Renderer.Render(listDeploymentsReply)
	return nil
}

type ListDeploymentsReply struct {
	mapper *idmapper.Mapper
	value  *koyeb.ListDeploymentsReply
	full   bool
}

func NewListDeploymentsReply(mapper *idmapper.Mapper, value *koyeb.ListDeploymentsReply, full bool) *ListDeploymentsReply {
	return &ListDeploymentsReply{
		mapper: mapper,
		value:  value,
		full:   full,
	}
}

func (ListDeploymentsReply) Title() string {
	return "Deployments"
}

func (r *ListDeploymentsReply) MarshalBinary() ([]byte, error) {
	return r.value.MarshalJSON()
}

func (r *ListDeploymentsReply) Headers() []string {
	return []string{"id", "service", "type", "status", "messages", "regions", "created_at"}
}

func (r *ListDeploymentsReply) Fields() []map[string]string {
	items := r.value.GetDeployments()
	resp := make([]map[string]string, 0, len(items))
	maxMessagesLength := 80
	if r.full {
		maxMessagesLength = 0
	}

	for _, item := range items {
		if item.GetStatus() == koyeb.DEPLOYMENTSTATUS_STASHED {
			continue
		}

		fields := map[string]string{
			"id":         renderer.FormatID(item.GetId(), r.full),
			"service":    renderer.FormatServiceSlug(r.mapper, item.GetServiceId(), r.full),
			"type":       formatDeploymentType(item.Definition.GetType()),
			"status":     formatDeploymentStatus(item.GetStatus()),
			"messages":   formatDeploymentMessages(item.GetMessages(), maxMessagesLength),
			"regions":    renderRegions(item.Definition.Regions),
			"created_at": renderer.FormatTime(item.GetCreatedAt()),
		}
		resp = append(resp, fields)
	}

	return resp
}
