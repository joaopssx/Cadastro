# sistema-cadastral

CLI em Go para gerenciar cadastros e salvar tudo num JSON.

## Como rodar

Precisa do Go 1.21+.

```bash
git clone https://github.com/usuario/sistema-cadastral.git
cd sistema-cadastral
go run ./src
```

## Estrutura

- `src/tipos.go`: structs do banco
- `src/storage.go`: lê e salva o arquivo json dentro da pasta db/
- `src/cadastro.go`: regras de negócio (inserir, remover, buscar, etc)
- `src/id.go`: gerador de ids aleatórios com 8 dígitos numéricos
- `src/main.go`: input do user e menu interativo no terminal
- `db/`: pasta onde fica armazenado o dados.json

## Funcionalidades

- Lista todos os cadastros
- Busca pelo ID
- Adiciona um novo (gera ID automaticamente e valida se o status é "pago" ou "atrasado")
- Remove pelo ID
- Altera a situação de um cadastro existente
- Mostra totais e um gráfico simples de barras usando `#`

## Menu

```text

Menu:
1 - listar todos
2 - buscar por ID
3 - adicionar
4 - remover
5 - alterar situacao
6 - estatisticas
0 - sair
Opcao: 
```
