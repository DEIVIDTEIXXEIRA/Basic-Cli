package killanddelete

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/process"
	"github.com/urfave/cli"
)

// KillAction finaliza os processos por meio de Id ou nome
func KillAction(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("nenhum argumento esperado, use flags")
	}

	if c.IsSet("id") && c.IsSet("name") {
		return errors.New("deve ser fornecida a flag para id ou name")
	}

	if !c.IsSet("id") && c.String("name") == "" {
		return errors.New("a flag name não pode estar vazia")
	}

	if erro := killProcess(c); erro != nil {
		return erro
	}
	fmt.Println("Processo finalizado com sucesso")
	return nil
}

func killProcess(c *cli.Context) error {
	// Encerra caso um id seja fornecido
	if c.IsSet("id") {
		processos, err := process.NewProcess(int32(c.Uint("id")))
		if err != nil {
			return err
		}

		return processos.Kill()
	}

	processos, err := process.Processes()
	if err != nil {
		return err
	}

	var (
		errs  []string
		found bool
	)

	target := c.String("name")
	for _, p := range processos {
		nome, _ := p.Name()
		if nome == "" {
			continue
		}

		if isEqualProcessName(nome, target) {
			found = true
			if err := p.Kill(); err != nil {
				e := err.Error()
				errs = append(errs, e)
			}
		}
	}

	if !found {
		return errors.New("processo não encontrado")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))

}


func isEqualProcessName(proc1 string, proc2 string) bool {
	// Verifica se dois nomes de processos são iguais
	if runtime.GOOS == "linux" {
		return proc1 == proc2
	}
	return strings.EqualFold(proc1, proc2)
}
