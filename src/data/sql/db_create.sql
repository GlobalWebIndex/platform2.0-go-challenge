
CREATE TABLE IF NOT EXISTS users (
    id integer primary key autoincrement,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    full_name varchar(255)
);
CREATE UNIQUE INDEX IF NOT EXISTS uix_users_username ON users(username);


CREATE TABLE IF NOT EXISTS audiences (
    id integer primary key autoincrement,
    title varchar(255) NOT NULL UNIQUE,
    birth_country varchar(255),
    gender integer,
    age_group_upper_limit integer,
    age_group_lower_limit integer
);
CREATE UNIQUE INDEX IF NOT EXISTS uix_audiences_title ON audiences(title);

CREATE TABLE IF NOT EXISTS audience_infos (
    id integer primary key autoincrement,
    audience_id integer,
    audience_info_type integer,
    audience_info_stat integer,
    FOREIGN KEY(audience_id) REFERENCES audiences(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS insights (
    id integer primary key autoincrement,
    title varchar(255) NOT NULL UNIQUE,
    insight_text varchar(255)
);
CREATE UNIQUE INDEX IF NOT EXISTS uix_insights_title ON insights(title);

CREATE TABLE IF NOT EXISTS charts (
    id integer primary key autoincrement,
    title varchar(255) NOT NULL UNIQUE,
    x_axes varchar(255),
    y_axes varchar(255)
);
CREATE UNIQUE INDEX IF NOT EXISTS uix_charts_title ON charts(title);

CREATE TABLE IF NOT EXISTS chart_points (
    id integer primary key autoincrement,
    chart_id integer,
    x real,
    y real,
    FOREIGN KEY(chart_id) REFERENCES charts(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS favorited_audiences (
    audience_id integer,
    user_id integer,
    favorited_description varchar(255),
    PRIMARY KEY (audience_id, user_id),
    FOREIGN KEY(audience_id) REFERENCES audiences(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    FOREIGN KEY(user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS favorited_charts (
    chart_id integer,
    user_id integer,
    favorited_description varchar(255),
    PRIMARY KEY (chart_id,user_id),
    FOREIGN KEY(chart_id) REFERENCES charts(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    FOREIGN KEY(user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS favorited_insights (
    insight_id integer,
    user_id integer,
    favorited_description varchar(255),
    PRIMARY KEY (insight_id, user_id),
    FOREIGN KEY(insight_id) REFERENCES insights(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    FOREIGN KEY(user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);
