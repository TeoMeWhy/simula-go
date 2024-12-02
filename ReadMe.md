# Simula Go

Uma experiência de simular álbuns de figurinhas utilizando a linguagem GoLang.

## Ideia

Como motivação de aprendizado, queremos entender como podemos trabalhar com processamento concorrente ao completar um grande volume de álbuns, minimizando assim o tempo de processo.

Para isso, utilizamos canais e Goroutines. Confira a implementação.

## Uso

Para executar o programa:

1. Tenha o Go instalado;
2. ```bash
cd src/figurinhas
go run main.go
```

Como padrão ele será executado com esses parâmetros:
```
  -N int
        Tamanho do álbum, i.e. quantas figurinhas são necessárias para completar o álbum. (default 550)
  -ite int
        Quantidade de albuns gerados na simulação. (default 1000)
  -n int
        Tamanho de cada pacote de figurinhas, i.e. quantas figurinahs em cada pacote. (default 5)
  -threads int
        Quantidade de processos na simulação. (default 1)
```

Ou seja, se preferir pode alter os valores para álbuns personalizados, bem como a quantidade de iteração e threads para procesamento.