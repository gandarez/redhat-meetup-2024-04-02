# Red Hat & Golang SP - Meetup 02/04/2024

A Red Hat em parceria com o Golang SP os convidam para o meetup em 02/04/2024.

## Desafio

O desafio é publicar uma API escrita em Golang no Red Hat OpenShift. Ao final sua API será capaz de ler e escrever em banco de dados e também consumir uma API externa.

## Desafio 1

Para rever o desafio 1, [clique aqui](/docs/CHALLENGE_1.md).

## Desafio 2

Para rever o desafio 2, [clique aqui](/docs/CHALLENGE_2.md).

## Desafio 3
![diagram](/assets/diagram/desafio03.jpg)

---
Vamos criar um banco de dados Redis no OpenShift que será utilizado para armazenar o token de autenticação da API externa. Em seguida vamos configurar as variáveis de ambiente para que a aplicação possa se conectar no Twitch e no IGDB.

- Você já aprendeu a criar um banco de dados no OpenShift, então crie um banco de dados Redis (Ephemeral). Defina a senha como **"redis"**.
- Agora vamos adicionar a secret que vai guardar o client id e o client secret do Twitch.
- No menu lateral clique em **"Secrets"**.
- Clique em **"Create"** e selecione **"Key/value secret"**.
- Defina o nome da secret como **"twitch-secret"**.
- Adicione as chaves **"client-id"** e **"client-secret"** com os valores do seu client id e client secret do Twitch.
- Agora vamos adicionar as variáveis de ambiente necessárias para a aplicação se conectar no Redis, Twitch e IGDB. Caso precise de ajuda, [clique aqui](https://api-docs.igdb.com/#getting-started) e siga as instruções em **"Account Creation"**.

  > Atenção especial as variáveis `VENDOR_TWITCH_CLIENT_ID` e `VENDOR_TWITCH_CLIENT_SECRET` que devem ser configuradas com os valores da secret que você criou.

    | Name | Value |
    | ---- | ----- |
    | REDIS_HOST | redis |
    | REDIS_PASSWORD | redis |
    | VENDOR_IGDB_HOST | <https://api.igdb.com/v4> |
    | VENDOR_TWITCH_HOST | <https://id.twitch.tv> |
    | VENDOR_TWITCH_CLIENT_ID | **"seu client id"** |
    | VENDOR_TWITCH_CLIENT_SECRET | **"seu client secret"** |

    ![env from secrets](/assets/screenshots/Screenshot_add_env_from_secret.png)

- Agora vamos testar o último endpoint da API que foi criado neste desafio. Para isso, você deve fazer uma requisição [GET] para o endpoint `/games` passando o nome do jogo como parâmetro. Exemplo: `/games/Mario`. A aplicação deve retornar uma lista de jogos que contem **"Mario"**, limitado a 10 resultados.

## Desafio Final :star: :star: :star: :star: :star:

No desafio final vamos testar a escalabilidade da aplicação. Para isso precisamos executar o **"Apache HTTP server benchmarking tool"** (ab) para simular uma carga de 10000 requisições concorrentes. O comando para executar o teste é:

```bash
ab -n 10000 -c 200 -k http://<URL>/games/Mario
```

> Substitua `<URL>` pela URL da sua aplicação.

- O teste deve ser executado com sucesso e a aplicação deve responder a todas as requisições sem falhar.
- A aplicação deve ser capaz de escalar horizontalmente para atender a demanda.
- Para verificar a escalabilidade, vá em **"Administrator"** > **"Workloads"** > **"Pods"**.
- Agora filtre por **"video-game-api"** e marque a opção **"Running"** para listar apenas os pods da aplicação.
- Você deve ver que a quantidade de pods aumentou para atender a demanda. A ideia aqui é que exista mais de um pod rodando simultaneamente até o limite de 5 concorrentes. A criação dos pods pode levar um tempo, então aguarde até que todos os pods estejam rodando.

## Parabéns :feelsgood:

Você concluiu o desafio. Agora você tem uma aplicação rodando no Red Hat OpenShift que é capaz de se conectar em banco de dados, consumir uma API externa e escalar horizontalmente para atender a demanda.

Made with :heart: by Red Hat and Golang SP
