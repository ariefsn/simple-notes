# Notes App

> Simple app for notes, created with Go & Vue.

## A. Stacks
```
- Knot 2.0
- Dbflex
- VueJS
- Vuex
- Viper
- MongoDB
- NPM or Yarn
```

## B. How To
1. Clone
```
# clone the project
cd $GOPATH/src
git clone "https://github.com/ayiexz/simple-notes.git"

# copy configuration file
cd $GOPATH/src/notes/app/config
cp config.json.template config.json
# adjust database config and other (if needed)
```

2. Run App
Set the value in configuration file to `dev` if want to development mode, and `prod` for production mode.

>#### Development Mode
> Need to run vue & main go.
```
# run the vue
cd app/views
npm run serve
or
yarn serve

# run main go
cd app
go run main.go
```

>#### Production Mode
> Build vue and run main go.
```
# build vue
cd app/views
npm run build
or
yarn build

# run main go
cd app
go run main.go
```

3. Open App
> Open App @ [http://127.0.0.1:8888/](http://127.0.0.1:8888/)

## C. Author
> 2019, Arief Setiyo Nugroho

## D. Thanks to
> [Noval Agung Prayogo](https://github.com/novalagung)