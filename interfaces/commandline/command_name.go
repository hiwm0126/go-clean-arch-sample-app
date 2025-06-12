package commandline

// CommandName コマンド名の型定義（型安全性向上のため）
type CommandName string

// コマンド名定数
const (
	CommandNameOrder    CommandName = "ORDER"
	CommandNameCancel   CommandName = "CANCEL"
	CommandNameShip     CommandName = "SHIP"
	CommandNameChange   CommandName = "CHANGE"
	CommandNameExpand   CommandName = "EXPAND"
	CommandNameInitData CommandName = "INIT_DATA"
)

// String CommandNameを文字列に変換
func (c CommandName) String() string {
	return string(c)
}

// IsValid コマンド名が有効かチェック
func (c CommandName) IsValid() bool {
	switch c {
	case CommandNameOrder, CommandNameCancel, CommandNameShip, CommandNameChange, CommandNameExpand, CommandNameInitData:
		return true
	default:
		return false
	}
}

// GetValidCommandNames 有効なコマンド名一覧を取得
func GetValidCommandNames() []CommandName {
	return []CommandName{
		CommandNameOrder,
		CommandNameCancel,
		CommandNameShip,
		CommandNameChange,
		CommandNameExpand,
		CommandNameInitData,
	}
} 