CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
COMMENT ON TABLE users IS 'Таблица для хранения информации о пользователях';
COMMENT ON COLUMN users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN users.name IS 'Имя пользователя';

CREATE TABLE breweries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
COMMENT ON TABLE breweries IS 'Таблица для хранения информации о пивоварнях';
COMMENT ON COLUMN breweries.id IS 'Уникальный идентификатор пивоварни';
COMMENT ON COLUMN breweries.name IS 'Название пивоварни';

CREATE TABLE beers_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    brewery_id INTEGER NOT NULL,
    CONSTRAINT fk_brewery
        FOREIGN KEY(brewery_id) REFERENCES breweries(id) ON DELETE CASCADE
);
COMMENT ON TABLE beers_types IS 'Таблица для хранения информации о сортах пива';
COMMENT ON COLUMN beers_types.id IS 'Уникальный идентификатор сорта пива';
COMMENT ON COLUMN beers_types.name IS 'Название сорта пива';
COMMENT ON COLUMN beers_types.brewery_id IS 'Идентификатор пивоварни, выпускающей этот сорт';

CREATE TABLE user_beer_ratings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    beer_id INTEGER NOT NULL,
    rating SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_beer
        FOREIGN KEY(beer_id) REFERENCES beers_types(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_beer_rating UNIQUE (user_id, beer_id)
);
COMMENT ON TABLE user_beer_ratings IS 'Таблица для хранения оценок сортов пива пользователями';
COMMENT ON COLUMN user_beer_ratings.id IS 'Уникальный идентификатор оценки';
COMMENT ON COLUMN user_beer_ratings.user_id IS 'Идентификатор пользователя, поставившего оценку';
COMMENT ON COLUMN user_beer_ratings.beer_id IS 'Идентификатор сорта пива, который оценили';
COMMENT ON COLUMN user_beer_ratings.rating IS 'Оценка от 1 до 5';
COMMENT ON CONSTRAINT unique_user_beer_rating ON user_beer_ratings IS 'Обеспечивает уникальность оценки пользователя для конкретного сорта пива';

