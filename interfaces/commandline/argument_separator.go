package commandline

// SeparatedArgument 分離されたコマンド（ArgumentSeparatorの生成物）
type SeparatedArgument struct {
	CommandName CommandName
	Args        [][]string
}

// ArgumentSeparator 引数分離の責任を持つインターフェース
type ArgumentSeparator interface {
	Separate(rawArgs [][]string) ([]SeparatedArgument, error)
}

// argumentSeparator 引数分離の実装
type argumentSeparator struct {
	validCommands map[CommandName]bool
}

// NewArgumentSeparator 引数分離器のコンストラクタ
func NewArgumentSeparator() ArgumentSeparator {
	return &argumentSeparator{
		validCommands: map[CommandName]bool{
			CommandNameOrder:    true,
			CommandNameCancel:   true,
			CommandNameShip:     true,
			CommandNameChange:   true,
			CommandNameExpand:   true,
			CommandNameInitData: true,
		},
	}
}

// Separate 生の引数をコマンド別に分離
func (p *argumentSeparator) Separate(rawArgs [][]string) ([]SeparatedArgument, error) {
	var separatedCommands []SeparatedArgument
	var currentArgs [][]string
	// 一個目のコマンド名が渡ってくるまでの値は、システム全体で利用される初期化用データ
	// そのため、最初のコマンド名はINIT_DATAで固定される
	var currentCommandName = "INIT_DATA"

	for _, arg := range rawArgs {
		if len(arg) == 0 {
			continue // 空のスライスはスキップ
		}

		// コマンド名が含まれているかチェック
		if p.isValidCommand(arg[0]) {
			// 前のコマンドを保存
			if err := p.appendSeparatedCommand(&separatedCommands, currentCommandName, currentArgs); err != nil {
				return nil, err
			}

			// 新しいコマンドに更新
			currentCommandName = arg[0]
			currentArgs = nil
		}
		currentArgs = append(currentArgs, arg)
	}

	// 最後のコマンドを保存
	if err := p.appendSeparatedCommand(&separatedCommands, currentCommandName, currentArgs); err != nil {
		return nil, err
	}

	return separatedCommands, nil
}

// isValidCommand コマンド名が有効かチェック
func (p *argumentSeparator) isValidCommand(commandName string) bool {
	_, exists := p.validCommands[CommandName(commandName)]
	return exists
}

// appendSeparatedCommand 分離されたコマンドをリストに追加
func (p *argumentSeparator) appendSeparatedCommand(separatedCommands *[]SeparatedArgument, commandName string, args [][]string) error {
	if len(args) == 0 {
		return nil
	}

	*separatedCommands = append(*separatedCommands, SeparatedArgument{
		CommandName: CommandName(commandName),
		Args:        args,
	})

	return nil
}
