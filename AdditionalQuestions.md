# Additional Questions

### 1. Are there any suggestions you could make about the structure of our data?

Some suggestions I would make:

> A. Similar to the structure you have with Genre table, you could create an Artists table, so in Songs table Artist attribute should be an int id.
Creating this kind of structure allows a fast obtaining of data when making a query and it also allows easier and dynamic modification.
In addition, it allows to have a normalized database. Creating catalogs are good habit for database design.

>B. If you want to have a more complete structure, point 'A.' is not enough. There are songs which could be featured by more than one artist.
In that case, it would be a better option to create an ArtistbySong table, which will store n artists by song using each own id.

> C. I would add more information about songs, just like album (using recommended structure in 'A.' above) and release date for example.

> D. I would rename song attribute in Songs table for 'title', I think it is clearer and more specific.

> E. I think artist size (varchar 1024), wherever this attribute is located, it could be a little bit smaller, perhaps varchar (256) would be enough.

### 2. What fields would you index in these tables?

> A. Artist id, if you made the recommended change, could be a good candidate for indexing, taking in consideration a huge amount of records by different artists.

> B. Genre id, could be a good candidate for indexing, taking in consideration a huge amount of records for songs and you wanted to query by genre.

> C. If id in Songs table is the primary key and identity, it will be enough.
