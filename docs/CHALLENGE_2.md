# Desafio 2

![diagram](/assets/diagrams/diagram_2.jpeg)

- Agora que a **"feature-2"** já faz parte da branch `master`, vamos testar os dois novos endpoints **"[GET] /consoles"** e **"[POST] /consoles"**.
- Se você tem **POSTMAN** isntalado siga essas [instruções](/docs/POSTMAN.md), caso contrário vamos testar usando **curl**.

> Não se esqueça de substituir `localhost:17020` pelo endereço da sua aplicação no OpenShift.

- **[GET] /consoles**

    ```bash
    curl --location 'http://localhost:17020/consoles/b171ae30-2d02-4da2-98b4-33ad2c331669' \
    --header 'Content-Type: application/json'
    ```

- Oops! Parece que eu me enganei, você recebeu um erro e sua aplicação está offline? Isso se dá ao fato de que a aplicação está tentando acessar um banco de dados que não existe. Vamos corrigir isso.
- Vamos criar um banco de dados PostgreSQL no OpenShift.
- No menu lateral clicar em **"+Add"** e escolher a opção **"Database"**.

    ![add database](/assets/screenshots/Screenshot_add_database.png)

- Escolher **"PostgreSQL (Ephemeral)"**.

    > Ephemeral significa que o banco de dados será apagado quando a aplicação for apagada.

    ![add postgresql](/assets/screenshots/Screenshot_add_postgresql.png)

  - Preencher o campo **"PostgreSQL Connection Username"** com **"postgresql"**.
  - Preencher o campo **"PostgreSQL Connection Password"** com **"postgresql"**.
  - Preencher o campo **"PostgreSQL Database Name"** com **"video-game-db"**.
  - Clicar em **"Create"** e aguardar a criação do banco de dados.
- Agora que o banco de dados foi criado, vamos adicionar as variáveis de ambiente necessárias para a aplicação acessar o banco de dados.
- No menu lateral clicar em **"Administrator"**, depois em **"Workloads"** e em **"Deployments"**.
- Clicar no nome da aplicação **"video-game-api"**.
- Clicar na aba **"Environment"**.
- Adicionar as seguintes variáveis de ambiente:

    | Name | Value |
    | ---- | ----- |
    | DATABASE_HOST | postgresql |
    | DATABASE_USER | postgresql |
    | DATABASE_PASSWORD | postgresql |
    | DATABASE_NAME | video-game-db |

- Clicar em **"Save"**.
- Aguarde a aplicação ser reiniciada.
- Agora que o banco de dados foi criado e as variáveis de ambiente foram adicionadas, vamos testar novamente os endpoints.

- **[GET] /consoles**

    ```bash
    curl --location 'http://localhost:17020/consoles/b171ae30-2d02-4da2-98b4-33ad2c331669' \
    --header 'Content-Type: application/json'
    ```

    Você deve receber uma resposta parecida com essa:

    ```json
    {
        "id": "b171ae30-2d02-4da2-98b4-33ad2c331669",
        "name": "Xbox 360",
        "manufacturer": "Microsoft",
        "release_date": "2005-11-22"
    }
    ```

- **[POST] /consoles**

    ```bash
    curl --location 'http://localhost:17020/consoles' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Super Nintendo",
        "manufacturer": "Nintendo",
        "release_date": "1990-11-21"
    }'
    ```

    Você deve receber uma resposta parecida com essa:

    ```json
    {
       "id": "65a87fdb-7f02-49db-a044-ab1bde71f5b0",
        "name": "Super Nintendo",
        "manufacturer": "Nintendo",
        "release_date": "1990-11-21"
    }

- Agora que já testamos os endpoints, vamos configurar o Horizontal Pod Autoscaler (*HPA*) para que a aplicação possa escalar automaticamente.

    > O Horizontal Pod Autoscaler (HPA) ajusta o número de pods em um deployment, replicaset ou statefulset.

- No menu lateral clicar em **"Administrator"**, depois em **"Workloads"** e em **"Deployments"**.
- Clicar nos três pontinhos e em **"Edit resource limits"**.

    > Os limites de recursos são usados para limitar a quantidade de recursos que um contêiner pode usar.

- Adicionar os seguintes limites de recursos:

    | Resource | Request | Limit |
    | -------- | ----- | ----- |
    | CPU | 50m| 100m |
    | Memory | 25Mi | 50Mi |

- Agora vamos configurar o HPA. Em topologia, selecione o **"video-game-api-app"** e no menu lateral que se abre, clique no menu superior direito em **"Actions"** e depois em **"Add HorizontalPodAutoscaler"**.
  - Preencher o campo **"name"** com **"hpa-video-game-api"**.
  - Preencher o campo **"Min Pods"** com **1**.
  - Preencher o campo **"Max Pods"** com **5**.
  - Preencher o campo **"CPU Utilization Target"** com **25%**.
  - Preencher o campo **"Memory Utilization Target"** com **50%**.
  - Clicar em **"Save"**.
