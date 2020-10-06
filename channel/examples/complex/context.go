// - Se a gente lança 500 goroutines pra fazer uma tarefa, e cancelamos a tarefa no meio do caminho, como fazemos pra matar as goroutines?
// - Documentação: https://golang.org/pkg/context/
// - Aos aventureiros: https://blog.golang.org/context
// - Destaques:
//     - ctx := context.Background
//     - ctx, cancel = context.WithCancel(context.Background)
//     - goroutine: select case ←ctx.Done(): return; default: continua trabalhando.
//     - check ctx.Err();
//     - Tambem tem WithDeadline/Timeout
// - Exemplos (Todd):
//     - Analisando:
//         - Background: https://play.golang.org/p/cByXyrxXUf
//         - WithCancel: https://play.golang.org/p/XOknf0aSpx
//         - Função Cancel: https://play.golang.org/p/UzQxxhn_fm
//     - Exemplos práticos:
//         - func WithCancel: https://play.golang.org/p/Lmbyn7bO7e
//         - func WithCancel: https://play.golang.org/p/wvGmvMzIMW
//         - func WithDeadline: https://play.golang.org/p/Q6mVdQqYTt
//         - func WithTimeout: https://play.golang.org/p/OuES9sP_yX
//         - func WithValue: https://play.golang.org/p/8JDCGk1K4P 