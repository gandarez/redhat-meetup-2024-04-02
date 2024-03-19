# Desafio 1

![diagram](/assets/diagrams/diagram_1.jpeg)

- Fazer fork neste [repositório](https://github.com/gandarez/redhat-meetup-2024-04-02).

  > Deixar desmarcada a opção **"Copy the master branch only"**, para indicar ao GitHub copiar todas as branches no fork.

  ![fork](/assets/screenshots/Screenshot_fork.png)

- Acessar sua conta no [Sandbox do OpenShift](https://console.redhat.com/openshift/sandbox).
- Escolher opção **"Developer Sandbox"** no menu lateral e clicar em **"Launch"** dentro do Red Hat OpenShift.

    ![launch sandbox](/assets/screenshots/Screenshot_launch_sandbox.png)

- No menu lateral clicar em **"+Add"** e escolher a opção **"Import from Git"**.

    ![import from git](/assets/screenshots/Screenshot_import_from_git.png)

- Preencher a url do Git com o **endereço do seu repositório** onde foi feito o fork.

    ![import from git](/assets/screenshots/Screenshot_import_from_git_2.png)

- Em Application Name, preencher com **"video-game-api"**.
- Em Name, preencher com **"video-game-api"** novamente.
- Em Resource Type, escolher **"Deployment"**.
- Somente marque a opção **"Add Pipeline"** se o time de organização do evento solicitar, caso contrário mantenha desabilitado.

> O OpenShift vai criar um pipeline para fazer o build e deploy da aplicação. O pipeline é uma sequência de passos que o OpenShift executa para construir a aplicação e publicar no ambiente.
>
> Com a opção desmarcada o OpenShift vai criar um Build automaticamente, porém sem um pipeline.

- Em Target Port preencher com **"17020"**.
- Na última linha clicar em **"Health checks"**.

  > Agora vamos configurar o Health Check da nossa API. O Health Check é uma forma de verificar se a aplicação está funcionando corretamente. O OpenShift faz isso verificando se a aplicação responde a uma requisição HTTP em um determinado caminho e porta. Se a aplicação responder corretamente, o OpenShift considera a aplicação saudável e pronta para receber tráfego. Se a aplicação não responder corretamente, o OpenShift considera a aplicação não saudável e não envia tráfego para ela.

  - Cliar em **"Add Readiness probe"**.
    - Em Path preencher com **"/readiness"**.
    - Em Port preencher com **"17020"**.
    - No canto inferior direito desse quadro clicar no check para confirmar.

    ![readiness probe](/assets/screenshots/Screenshot_readiness_probe.png)

  - Clicar em **"Add Liveness probe"**.
    - Em Path preencher com **"/liveness"**.
    - Em Port preencher com **"17020"**.
    - No canto inferior direito desse quadro clicar no check para confirmar.

    ![liveness probe](/assets/screenshots/Screenshot_liveness_probe.png)

- Ainda na última linha clicar em **"Deployment"**.

  > Deixar a opção "Auto Deploy" marcada. Isso faz com que o OpenShift faça o deploy da aplicação automaticamente quando o build terminar.

  - Agora vamos definir as variáveis de ambiente necessárias nesse primeiro desafio.
  - Insira "SERVICE_NAME" no campo **"Name"** e "video-game-api" no campo **"Value"**.

- Clicar em **"Create"**.

    > O OpenShift vai criar um pipeline para fazer o build e deploy da aplicação. O pipeline é uma sequência de passos que o OpenShift executa para construir a aplicação e publicar no ambiente.

- Agora você foi redirecionado para a página da topologia.

  ![topology](/assets/screenshots/Screenshot_topology.png)

- Aguarde o término da pipeline. Caso queira acompanhar o progresso, clique em **"Pipelines"** no menu lateral e espere ele ficar verde.

    ![pipeline running](/assets/screenshots/Screenshot_pipeline_running.png)

- Quando o pipeline terminar, a sua API estará demarcada com um circulo azul conforme imagem acima na topologia.
- Agora vamos testar o seu primeiro endpoint em Produção. Clique no botão redondo no canto superior direito do circulo azul.

  ![open url](/assets/screenshots/Screenshot_open_url.png)

- No seu browser vai abrir uma nova aba com a URL da sua API. Adicione **"/readiness"** ou **"/liveness"** no final da URL e pressione Enter.

  > Você deve ver um Ok na tela. Isso significa que a sua aplicação está saudável e pronta para receber tráfego.

## Extra - Desafio Pipeline

**Este passo deve ser executado somente se o pipeline foi adicionado no momento da criação da aplicação.**

Para nossa API ser publicada automaticamente, precisamos adicionar um webhook no GitHub para notificar o OpenShift quando houver alterações no repositório. Vamos considerar o cenário mais simples, quando houver um push na branch `master`.

- No menu lateral clique em **"Pipelines"**.
- Na coluna Name clique no nome da sua aplicação **"video-game-api"**.
- No canto superior direito em Actions clique em **"Add Trigger"**.

  ![add trigger](/assets/screenshots/Screenshot_add_trigger.png)

- Em Git Provider Type procure por **"github-push"**.
- Em Git Revision digite **"master"**.
- Clique em **"Add"**.

  > Se você receber o erro **"admission webhook "webhook.triggers.tekton.dev" denied the request: mutation failed: cannot decode incoming new object: json: unknown field "name""** você pode pular essa etapa, porém você terá que iniciar os builds manualmente.

## Extra - Desafio Build

**Este passo deve ser executado somente se o pipeline NÃO foi adicionado no momento da criação da aplicação.**

Para nossa API ser publicada automaticamente, precisamos adicionar um webhook no GitHub para notificar o OpenShift quando houver alterações no repositório. Vamos considerar o cenário mais simples, quando houver um push na branch `master`.

- No menu lateral clique em **"Builds"**.
- Na coluna Name clique no nome da sua aplicação **"video-game-api"**.
- Na última linha em Webhooks clique em **"Copy URL with Secret"** na linha do GitHub.
- Vá até seu repositório no GitHub e clique em **"Settings"**.
- No menu lateral clique em **"Webhooks"**.
- Clique em **"Add webhook"**.
- Em Payload URL cole o URL que você copiou do OpenShift.
- Em Content type selecione **"application/json"**.
- Em "Which events would you like to trigger this webhook?" selecione **"Just the push event"**.
- Clique em **"Add webhook"**.

  > Agora toda vez que você fizer um push na branch `master` do seu repositório, o OpenShift vai iniciar um novo build da sua aplicação.
