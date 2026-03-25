package commandline

import (
	"strings"
	"testing"

	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

func TestValidateInitDataQueryCount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		cmds    []cli.ParsedCommand
		wantErr string
	}{
		{
			name: "空ならOK",
			cmds: nil,
		},
		{
			name: "先頭がINIT_DATAでないならOK",
			cmds: []cli.ParsedCommand{{Kind: cmdname.CommandNameOrder, Order: &usecase.OrderUseCaseReq{}}},
		},
		{
			name: "NumOfQueryと後続件数が一致",
			cmds: []cli.ParsedCommand{
				{Kind: cmdname.CommandNameInitData, InitData: &usecase.DataInitializationUseCaseReq{NumOfQuery: 2}},
				{Kind: cmdname.CommandNameOrder, Order: &usecase.OrderUseCaseReq{}},
				{Kind: cmdname.CommandNameCancel, Cancel: &usecase.CancelUseCaseReq{}},
			},
		},
		{
			name: "件数不一致ならエラー",
			cmds: []cli.ParsedCommand{
				{Kind: cmdname.CommandNameInitData, InitData: &usecase.DataInitializationUseCaseReq{NumOfQuery: 1}},
				{Kind: cmdname.CommandNameOrder, Order: &usecase.OrderUseCaseReq{}},
				{Kind: cmdname.CommandNameCancel, Cancel: &usecase.CancelUseCaseReq{}},
			},
			wantErr: "query count mismatch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validateInitDataQueryCount(tt.cmds)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}
