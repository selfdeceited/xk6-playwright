FROM golang:buster
RUN apt-get update
RUN apt install git
RUN git clone https://github.com/grafana/k6.git
RUN go install github.com/mxschmitt/playwright-go/cmd/playwright@latest
RUN playwright install --with-deps
RUN apt-get install -y \
    fonts-liberation \
    gconf-service \
    libappindicator1 \
    libasound2 \
    libatk1.0-0 \
    libcairo2 \
    libcups2 \
    libfontconfig1 \
    libgbm-dev \
    libgdk-pixbuf2.0-0 \
    libgtk-3-0 \
    libicu-dev \
    libjpeg-dev \
    libnspr4 \
    libnss3 \
    libpango-1.0-0 \
    libpangocairo-1.0-0 \
    libpng-dev \
    libx11-6 \
    libx11-xcb1 \
    libxcb1 \
    libxcomposite1 \
    libxcursor1 \
    libxdamage1 \
    libxext6 \
    libxfixes3 \
    libxi6 \
    libxrandr2 \
    libxrender1 \
    libxss1 \
    libxtst6 \
    xdg-utils
WORKDIR /go/k6
RUN CGO_ENABLED=0 go install -a -trimpath -ldflags "-s -w -X ./lib/consts.VersionDetails=$(date -u +"%FT%T%z")/$(git describe --always --long --dirty)"
RUN go install go.k6.io/xk6/cmd/xk6@latest
RUN xk6 build v0.36.0 --with github.com/mstoykov/xk6-counter --with github.com/szkiba/xk6-dotenv --with github.com/dgzlopes/xk6-notifications --with github.com/grafana/xk6-output-influxdb --with github.com/wosp-io/xk6-playwright
RUN cp k6 $GOPATH/bin/k6
WORKDIR /home/k6
ENTRYPOINT ["k6"]