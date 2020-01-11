# Google App Engine Single Page Application Project Base

This is my Project Base for a Single Page Application (SPA) built to use Google App Engine (GAE). 

Google App Engine is a quick and easy way to build and deploy application on a fully managed platform. It dynamically scales the application without having to manage the underlying infrastructure and is a great tool for prototyping to complete production site. The best part is because GAE automatically scales depending on application traffic and resource use, the cost is also scaled based on usage. [more reading](https://cloud.google.com/appengine/)

Single Page Application is a web application that dynamically updates the current page rather than loading entire new pages every time the user performs an action. All the necessary code (HTML, JS, CSS) are loaded initially in a single page load and then subsequent user actions dynamically loads appropriate additional resources to add to the page, never completely reloading the page at any point. [more reading](https://en.wikipedia.org/wiki/Single-page_application)

Project Base is my starting point for building any of my SPA. It contains all the basic setup and integration that I will need between the front page application and the back end server. I can load my project starter into my Integrated Development Environment (IDE) and just build it and it will run with a basic set of functionality and layout that I can then build on top of. It goes beyond just a builder plate code or starter kit; it's a standalone application ready for me to expand.

## Getting Started

These instructions will tell you how to deploy the project base to your IDE.

### Prerequisites

Because this is a fully working basic SPA, it will require an IDE with the following development tools already installed and working.

* Node.js (v12.14)
* NPM (v6.13)
* GoLang (v1.11)
* Angular (v8.3)
* Google Cloud SDK
* Google Cloud App Engine for Go
* Google Cloud Datastore Emulator

[PREREQUISITE.md](PREREQUISITE.md) provides instruction on how to setup an IDE using CodeAnywhere.

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