package visuailize

import (
	model "../ksql-server/model"
	util "../util"
	"fmt"
	"github.com/alexeyco/simpletable"
	_ "github.com/alexeyco/simpletable"
	"strconv"
)

func VisualizeStream(response []model.KsqlStreamResponse) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Topic"},
			{Align: simpletable.AlignCenter, Text: "Format"},
		},
	}

	for _, r := range response {
		for _, stream := range r.Streams {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: stream.Name},
				{Text: stream.Topic},
				{Align: simpletable.AlignRight, Text: stream.Format},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeTable(response []model.KsqlTableResponse) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Topic"},
			{Align: simpletable.AlignCenter, Text: "Format"},
		},
	}

	for _, r := range response {
		for _, tableDetail := range r.Tables {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignRight, Text: tableDetail.Name},
				{Text: tableDetail.Topic},
				{Align: simpletable.AlignRight, Text: tableDetail.Format},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeTopic(response []model.KsqlTopicResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},

			{Align: simpletable.AlignCenter, Text: "Replica Count"},
		},
	}

	for _, r := range response {
		for _, topic := range r.Topics {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: topic.Name},
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(len(topic.ReplicaInfo))},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeQueryList(response []model.KsqlQueryListResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
		},
	}

	for _, r := range response {
		for _, query := range r.Queries {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: query.Id},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeCrudStatementResponse(response []model.CrudStatementResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Statement"},
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "SequenceNumber"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Message"},
		},
	}

	for _, statement := range response {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: statement.StatementText},
			{Align: simpletable.AlignCenter, Text: statement.CommandID},
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(statement.CommandSequenceNumber)},
			{Align: simpletable.AlignCenter, Text: statement.CommandStatus.Status},
			{Align: simpletable.AlignCenter, Text: statement.CommandStatus.Message},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeDescribeResponse(response []model.DescribeResponse) {
	fields := simpletable.New()

	fields.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Field"},
			{Align: simpletable.AlignCenter, Text: "Type"},
		},
	}

	for _, detail := range response {
		for _, field := range detail.SourceDescription.Fields {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: field.Name},
				{Align: simpletable.AlignCenter, Text: field.Schema.Type},
			}
			fields.Body.Cells = append(fields.Body.Cells, row)

		}
	}

	fields.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(fields.String())

	fmt.Println("\n\n--------------------")

	fmt.Println("Run Time Statistics")
	for _, detail := range response {
		fmt.Println(detail.SourceDescription.Statistics)
	}
}

func VisualizeSelectStatementHeader(header []string, table *simpletable.Table) {

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{},
	}

	for _, s := range header {
		cell := simpletable.Cell{
			Align: simpletable.AlignCenter,
			Text:  s,
		}

		table.Header.Cells = append(
			table.Header.Cells,
			&cell,
		)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

}

func VisualizeSelectStatementBody(response model.SelectStatementBodyResponse, table *simpletable.Table) {
	var rows []*simpletable.Cell
	for _, column := range response.Row.Columns {

		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: util.ToString(column)},
		}

		rows = append(rows, row...)

	}

	table.Body.Cells = append(table.Body.Cells, rows)

	table.SetStyle(simpletable.StyleCompactLite)
	table.PrintLastInserted()
}

func VisualizeServerInfo(response model.KsqlServerInfoResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "version"},
			{Align: simpletable.AlignCenter, Text: "kafkaClusterId"},
			{Align: simpletable.AlignCenter, Text: "ksqlServiceId"},
		},
	}

	info := response.KsqlServerInfo
	row := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: info.Version},
		{Align: simpletable.AlignCenter, Text: info.KafkaClusterID},
		{Align: simpletable.AlignCenter, Text: info.KsqlServiceId},
	}

	table.Body.Cells = append(table.Body.Cells, row)

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizeServerHealthCheck(response model.KsqlServerHealthCheckResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "isHealthy"},
			{Align: simpletable.AlignCenter, Text: "isKafkaHealthy"},
		},
	}

	row := []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Text: strconv.FormatBool(response.IsHealthy)},
		{Align: simpletable.AlignCenter, Text: strconv.FormatBool(response.Details.Kafka.IsHealthy)},
	}

	table.Body.Cells = append(table.Body.Cells, row)

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func VisualizePropertyList(response []model.ListPropertiesResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Scope"},
			{Align: simpletable.AlignCenter, Text: "Values"},
		},
	}

	for _, r := range response {
		for _, query := range r.Properties {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: query.Name},
				{Align: simpletable.AlignLeft, Text: query.Scope},
				{Align: simpletable.AlignLeft, Text: query.Value},
			}
			table.Body.Cells = append(table.Body.Cells, row)
		}
	}
	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}
