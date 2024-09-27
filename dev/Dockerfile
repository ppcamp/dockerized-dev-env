FROM ubuntu:noble

# ARG USER

RUN apt update -y
RUN apt upgrade -y

RUN apt install -y sudo golang nodejs vim curl git

# Install vscode
RUN apt install software-properties-common apt-transport-https wget -y
RUN apt update -y
RUN wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main"
RUN apt install -y code

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/cweill/gotests/gotests@v1.6.0
RUN go install github.com/fatih/gomodifytags@v1.17.0
RUN go install github.com/josharian/impl@v1.4.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

RUN mkdir -p $HOME/app
WORKDIR $HOME/app

CMD ["bash"]