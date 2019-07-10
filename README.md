# Bible Go API

[![Build Status](https://travis-ci.org/rkeplin/bible-go-api.svg?branch=master)](https://travis-ci.org/rkeplin/bible-go-api)

Bible Go API is an open source REST API.  It contains multiple translations of The Holy Bible, as well as cross-references. 
This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY.

### Live Demo
A live demo of this application can be viewed [here](https://bible-go-api.rkeplin.com/v1/books/1/chapters/1).

### Getting Everything Running
```bash
git clone https://www.github.com/rkeplin/bible-go-api
cd bible-go-api && make images && make up
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

### Related Projects
* [Bible PHP API](https://www.github.com/rkeplin/bible-php-api)
* [Bible AngularJS UI](https://www.github.com/rkeplin/bible-angularjs-ui)
* [Bible MariaDB Docker Image](https://www.github.com/rkeplin/bible-mariadb)

### Credits
Data for this application was gathered from the following repositories.
* [scrollmaper/bible_database](https://github.com/scrollmapper/bible_databases)
* [honza/bibles](https://github.com/honza/bibles)

### License
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.
