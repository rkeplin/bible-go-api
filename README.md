# Bible Go API

[![Build Status](https://travis-ci.org/rkeplin/bible-go-api.svg?branch=master)](https://travis-ci.org/rkeplin/bible-go-api)

Bible Go API is an open source REST API.  It contains multiple translations of The Holy Bible, as well as cross-references. 
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY.

### Live Demo
A live demo of this application can be viewed [here](https://bible-go-api.rkeplin.com/v1/books/1/chapters/1).

### Getting Everything Running
```bash
git clone https://www.github.com/rkeplin/bible-go-api
cd bible-go-api
make images
make up
```
Note: Upon first start, the volume containing the MySQL data may take several seconds to load.

You should then be able to access [http://localhost:8084](http://localhost:8084) for the REST API and [http://localhost:8082](http://localhost:8082) for the UI (AngularJS).

### API Specifications
#### List of available translations
```bash
GET http://localhost:8084/translations
GET http://localhost:8084/translations/[TranslationID]
```

#### List of Genres
```bash
GET http://localhost:8084/genres
GET http://localhost:8084/genres/[GenreID]
```

#### Content
```bash
GET http://localhost:8084/books
GET http://localhost:8084/books/[BookID]
GET http://localhost:8084/books/[BookID]/chapters/[ChapterID]
GET http://localhost:8084/books/[BookID]/chapters/[ChapterID]
GET http://localhost:8084/books/[BookID]/chapters/[ChapterID]/[VerseID]
```
Note: In order to get content for a specific translation, supply `translation` as a Query Parameter.  For example,
`http://localhost:8084/books/1/chapters/1/1001002?translation=ASV`

#### Cross References
```bash
GET http://localhost:8084/verse/[VerseID]/relations 
```

#### Search
```bash
GET http://localhost:8084/search?query=[Query]
GET http://localhost:8084/searchAggregator?query=[Query]
```

### Related Projects
* [Bible Go API](https://www.github.com/rkeplin/bible-go-api)
* [Bible PHP API](https://www.github.com/rkeplin/bible-php-api)
* [Bible AngularJS UI](https://www.github.com/rkeplin/bible-angularjs-ui)
* [Bible MariaDB Docker Image](https://www.github.com/rkeplin/bible-mariadb)

### Credits
Data for this application was gathered from the following repositories.
* [scrollmaper/bible_database](https://github.com/scrollmapper/bible_databases)
* [honza/bibles](https://github.com/honza/bibles)
