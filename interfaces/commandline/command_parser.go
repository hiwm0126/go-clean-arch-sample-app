package commandline

import "errors"

// Parser 全体調整の責任を持つインターフェース
type Parser interface {
	Execute(rawArgs [][]string) ([]interface{}, error)
}

// commandParser 全体調整の実装
type commandParser struct {
	argumentSeparator ArgumentSeparator
	dataConverter     DataConverter
	parsers           []CommandArgumentParser
}

// NewParser コマンドパーサーのコンストラクタ
func NewParser() Parser {
	parser := &commandParser{
		argumentSeparator: NewArgumentSeparator(),
		dataConverter:     NewDataConverter(),
		parsers:           make([]CommandArgumentParser, 0),
	}

	// デフォルトハンドラーを登録
	parser.registerParser(NewOrderArgumentParser())
	parser.registerParser(NewCancelArgumentParser())
	parser.registerParser(NewShipArgumentParser())
	parser.registerParser(NewChangeArgumentParser())
	parser.registerParser(NewExpandArgumentParser())
	parser.registerParser(NewInitDataArgumentParser())

	return parser
}

// RegisterHandler パーサーを登録
func (p *commandParser) registerParser(handler CommandArgumentParser) {
	p.parsers = append(p.parsers, handler)
}

// Parse 引数をパースしてリクエストパラメータリストを返す
func (p *commandParser) Execute(rawArgs [][]string) ([]interface{}, error) {
	// 1. 引数をコマンド別に分離
	separatedCommands, err := p.argumentSeparator.Separate(rawArgs)
	if err != nil {
		return nil, err
	}

	// 2. 各コマンドを適切なハンドラーで処理
	var paramList []interface{}
	for _, separatedCommand := range separatedCommands {
		param, err := p.parseArgument(separatedCommand)
		if err != nil {
			return nil, err
		}
		paramList = append(paramList, param)
	}

	return paramList, nil
}

// parseArgument 個別のコマンドをパース
func (p *commandParser) parseArgument(separatedArgument SeparatedArgument) (interface{}, error) {
	// 適切なハンドラーを見つける
	for _, parser := range p.parsers {
		if parser.CanHandle(separatedArgument.CommandName) {
			return parser.Parse(separatedArgument.Args, p.dataConverter)
		}
	}

	return nil, errors.New("no parser found for command: " + string(separatedArgument.CommandName))
}
