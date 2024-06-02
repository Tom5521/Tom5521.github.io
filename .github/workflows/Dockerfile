FROM fedora:latest

RUN dnf install -y just wget

# Install golang
RUN wget -O go.tar.gz "https://go.dev/dl/go1.19.3.linux-amd64.tar.gz"
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz
ENV PATH="$PATH:/usr/local/go/bin"

# Install fyne package
RUN go install fyne.io/fyne/v2/cmd/fyne@latest

# Install gopherjs
RUN go install github.com/gopherjs/gopherjs@v1.19.0-beta1

COPY . .

RUN just build-js
