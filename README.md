# Liquid Night to Shine QR Code Reader

This is a Single Page application used for liquid church's night to shine event. It consist of a barcode scanner to help lookup guest information in an easy to use webapp.

## Getting Started

These instructions will tell you how to deploy the project base to your IDE.

### Prerequisites

For development, this application requires an IDE with the following development tools already installed and working.

* Node.js (v12.14)
* NPM (v6.13)
* GoLang (v1.11)
* Angular (v8.3)
* Google Cloud SDK
* Google Cloud App Engine for Go
* Google Cloud Datastore Emulator

### Installing

Clone the Starter Base into your project directory. Make sure your current directory is your project directory and it's empty.
```
$ git clone git@github.com:geoct826/gae-spa-base.git ./
```

Install dependent packages through NPM Postinstall script.
```
$ npm run postinstall
```

Build the webapp through npm build script.
```
$ npm run build
```

Start GAE development server
```
$ npm start
```

## Deployment

To deploy the SPA to Google Cloud, first you have to create a Project and enable [Google App Engine](https://console.cloud.google.com/appengine/start). This SPA is in Go lanaguage based on a Standard runtime. Once the Project and App Engine enabled, deploying the application is pretty easy

### GCloud Configuration

To start gcloud configuration, run 
```
$ gcloud init
```

This will take you through an OAuth2 authentication process to get the Google Cloud account linked to the IDE. Once gcloud initilization is complete, you should be able to deploy to google cloud by running
```
$ npm deploy
```
which runs gcloud app deploy.

## Built With

### Web Application
* [Angular](http://angular.io) - Web application framework
* [npm](https://npmjs.com/get-npm/) - Package and dependency management
* [Zurb Foundation](https://get.foundation/) - Responsive front-end framework
* [Apollo GraphQL](https://apollographql.com) - Query Language

### App Server
* [Go](https://golang.org) - Server programming language
* [Google App Engine](https://cloud.google.com/appengine/) - Fully managed platform
* [Google Cloud Datastore](https://cloud.google.com/datastore/) - NoSQL database

## Versioning

[SemVer](http://semver.org/) is used for versioning. For the versions available, see the [CHANGELOG.md](CHANGELOG.md). 

## Authors

* **George Tuan**


## License

This project is licensed under the MIT License.