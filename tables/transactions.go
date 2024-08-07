package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// func GetTransactionsTable(ctx *context.Context) table.Table {

// 	transactions := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

// 	info := transactions.GetInfo().HideFilterArea()

// 	info.SetTable("public.transactions").SetTitle("Transactions").SetDescription("Transactions")

// 	formList := transactions.GetForm()

// 	formList.SetTable("public.transactions").SetTitle("Transactions").SetDescription("Transactions")

// 	return transactions
// }

func GetTransactionsTable(ctx *context.Context) (transactionsTable table.Table) {
	transactionsTable = table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := transactionsTable.GetInfo().SetFilterFormLayout(form.LayoutFilter)
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("User ID", "user_id", db.Int).FieldFilterable()
	info.AddField("Amount", "amount", db.Decimal).FieldFilterable()
	info.AddField("Transaction Date", "transaction_date", db.Timestamp).FieldSortable()
	info.AddField("Description", "description", db.Text)

	// Add a custom button if needed, similar to the Authors example
	info.AddButton(ctx, "View Details", icon.Tv,
		action.PopUpWithIframe("/transactions/details", "Transaction Details", action.IframeData{Src: "/admin/info/transactions"}, "900px", "560px"))

	info.SetTable("transactions").SetTitle("Transactions").SetDescription("Transactions")

	formList := transactionsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("User ID", "user_id", db.Int, form.Number)
	formList.AddField("Amount", "amount", db.Decimal, form.Text)
	formList.AddField("Transaction Date", "transaction_date", db.Timestamp, form.Datetime)
	formList.AddField("Description", "description", db.Text, form.TextArea)

	formList.SetTable("transactions").SetTitle("Transactions").SetDescription("Transactions")

	return
}
