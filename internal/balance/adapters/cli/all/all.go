//nolint:revive // стандартный способ подключить подкоманды cobra к корневой
package all

import (
	_ "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/list_users"
	_ "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/signin"
	_ "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/signup"
)
