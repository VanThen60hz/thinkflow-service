package common

import (
	"github.com/VanThen60hz/service-context/core"
	"gorm.io/datatypes"
)

type SimpleTranscript struct {
	core.SQLModel
	Content string `json:"content"`
}

func (SimpleTranscript) TableName() string {
	return "transcripts"
}

func NewSimpleTranscript(id int, content string) SimpleTranscript {
	return SimpleTranscript{
		SQLModel: core.SQLModel{Id: id},
		Content:  content,
	}
}

type SimpleSummary struct {
	core.SQLModel
	SummaryText string `json:"summary_text"`
}

func (SimpleSummary) TableName() string {
	return "summaries"
}

func NewSimpleSummary(id int, summaryText string) SimpleSummary {
	return SimpleSummary{
		SQLModel:    core.SQLModel{Id: id},
		SummaryText: summaryText,
	}
}

type SimpleMindmap struct {
	core.SQLModel
	MindmapData datatypes.JSON `json:"mindmap_data"`
}

func (SimpleMindmap) TableName() string {
	return "mindmaps"
}

func NewSimpleMindmap(id int, mindmapData datatypes.JSON) SimpleMindmap {
	return SimpleMindmap{
		SQLModel:    core.SQLModel{Id: id},
		MindmapData: mindmapData,
	}
}
