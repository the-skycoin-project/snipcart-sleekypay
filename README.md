# SleekyPay
SleekyPay is a fake payment gateway to showcase Snipcart's custom payment gateway integration functionality.
[demo.snipcart.com](https://demo.snipcart.com)

## Local development

Netlify CLI must be installed globally

    npm install netlify-cli -g

To run the project locally, you can use the following commands:
```
npm install
npm update
npm run dev
```
The website will be live on `localhost:1234/index.html` while netlify functions can be found on `http://localhost:34567/.netlify/functions/<your_function_name>`

## Production build
```
npm run build
```

## Statik sources build

```
go get github.com/rakyll/statik/fs
sudo ln -s ~/go/bin/statik /usr/bin/statik
static -src=./dist
go generate
```

## Production Deploy

hosting the `/dist` folder should suffice for viewing the app.

A golang implementation of this has been provided which hosts the app on 127.0.0.1:8041

```
go mod vendor -v
go run main.go
```

a single binary can also be built with the GUI sources
```
go build .
```

I suggest reverse-proxying port 8041 to a subdomain of your website using caddy server
