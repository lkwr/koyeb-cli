package koyeb

import (
	"context"
	"fmt"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *DeploymentHandler) Get(cmd *cobra.Command, args []string) error {
	client := getApiClient()
	ctx := getAuth(context.Background())

	res, _, err := client.DeploymentsApi.GetDeployment(ctx, args[0]).Execute()
	if err != nil {
		fatalApiError(err)
	}

	full, _ := cmd.Flags().GetBool("full")
	getDeploymentsReply := NewGetDeploymentReply(&res, full)

	output, _ := cmd.Flags().GetString("output")
	return renderer.NewItemRenderer(getDeploymentsReply).Render(output)
}

type GetDeploymentReply struct {
	res  *koyeb.GetDeploymentReply
	full bool
}

func NewGetDeploymentReply(res *koyeb.GetDeploymentReply, full bool) *GetDeploymentReply {
	return &GetDeploymentReply{
		res:  res,
		full: full,
	}
}

func (a *GetDeploymentReply) MarshalBinary() ([]byte, error) {
	return a.res.GetDeployment().MarshalJSON()
}

func (a *GetDeploymentReply) Title() string {
	return "Deployment"
}

func (a *GetDeploymentReply) Headers() []string {
	return []string{"id", "app", "service", "status", "status_message", "created_at"}
}

func (a *GetDeploymentReply) Fields() []map[string]string {
	res := []map[string]string{}
	item := a.res.GetDeployment()
	fields := map[string]string{
		"id":             renderer.FormatID(item.GetId(), a.full),
		"app":            renderer.FormatID(item.GetAppId(), a.full),
		"service":        renderer.FormatID(item.GetServiceId(), a.full),
		"status":         formatDeploymentStatus(item.State.GetStatus()),
		"status_message": item.State.GetStatusMessage(),
		"created_at":     renderer.FormatTime(item.GetCreatedAt()),
	}
	res = append(res, fields)
	return res
}

func formatDeploymentStatus(ds koyeb.ServiceRevisionStateStatus) string {
	return fmt.Sprintf("%s", ds)
}