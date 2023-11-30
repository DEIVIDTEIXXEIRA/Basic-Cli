package main

import (
	killanddelete "basicCli/killAndDelete"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// Cria uma nova instância do aplicativo CLI
	app := cli.NewApp()
	app.Name = "Implementação básica de comando kill e delete CLI"
	app.Usage = "Permite que você encerre processos por nome ou ID e exclua arquivos ou pastas"

	// Define os comandos disponíveis no aplicativo
	app.Commands = []cli.Command{
		{
			Name:        "kill",
			HelpName:    "kill",
			Action:      killanddelete.KillAction,
			ArgsUsage:   ` `,
			Usage:       `encerra processos por ID de processo ou nome de processo.`,
			Description: `Encerra um processo.`,
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:  "id",
					Usage: "encerra o processo pelo ID do processo.",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "encerra o processo pelo nome do processo.",
				},
			},
		},
		{
			Name:        "volumes",
			HelpName:    "volumes",
			Action:      killanddelete.ActionVolumes,
			ArgsUsage:   `  `,
			Usage:       `lista os volumes do sistema de arquivos montados.`,
			Description: `Lista os volumes montados.`,
		},
	}

	// Executa o aplicativo CLI
	erro := app.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
