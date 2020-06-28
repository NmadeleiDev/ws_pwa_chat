const Sqlite = require("nativescript-sqlite");

const sqliteDB = {
    state: () => ({
        conn: {},
        data: []
    }),
    mutations: {
        INIT_SQLITE(state, data) {
            state.conn = data.database;
            console.log("Sqlite initialized successfully");
        },
    },
    actions: {
        async INIT_SQLITE_CONN(context) {
            let db;

            try {
                db = await new Sqlite("ws_chat.db")
            } catch (e) {
                console.log("Error opening db: ", e)
                return false
            }

            try {
                // await db.execSQL("DROP TABLE user_data")
                await db.execSQL("CREATE TABLE IF NOT EXISTS user_data (username TEXT UNIQUE, session_key TEXT DEFAULT '', user_secret TEXT, active INTEGER DEFAULT 1)")
                let existingData = await db.all("SELECT * FROM user_data")
                console.log("Db content: ", existingData);
            } catch (e) {
                console.log("Error creating table: ", e)
                return false
            }
            context.commit("INIT_SQLITE", { database: db });
            return true
        },
        async SAVE_KEY_TO_DB(context, payload) {
            let res;
            try {
                await context.state.conn.execSQL("UPDATE user_data SET session_key=? WHERE username=?", [[payload.key], [payload.username]]);
                console.log("Session key saved: ", payload)
                res = true
            } catch (e) {
                console.log("SQL ERROR", e);
                res = false
            }
            return res;
        },
        async SAVE_NEW_USER_TO_DB(context, user) {
            let insertQuery = "INSERT INTO user_data(username, user_secret) VALUES(?, ?)"
            let updateQuery = "UPDATE user_data SET user_secret=?, active=1 WHERE username=?"
            console.log("Created user secret: ", user.userSecret)
            try {
                await context.state.conn.execSQL("UPDATE user_data SET active=0"); // ставим всех существующих юзеров как неактивых, чтобы новый был единственным активным
                await context.state.conn.execSQL(insertQuery, [[user.username], [user.userSecret]]);
                return true
            } catch (e) {
                try { // да, это ужастный костыль, но это необходимо, так как в этом sqlite походу не работает ON CONFLICT
                    await context.state.conn.execSQL(updateQuery, [[user.userSecret], [user.username]]);
                    return true;
                } catch (e) {
                    console.log(e)
                }
                return false
            }
        },
        async LOAD_ACTIVE_USER_SECRET_KEYS(context) {
            let user;
            let data;
            try {
                data = await context.state.conn.get("SELECT username, session_key, user_secret FROM user_data WHERE active=1");
                if (data.length < 3 || data.find(item => item.length === 0) !== undefined) {
                    console.log("Failed to load user db secrets. Data: ", data);
                    return false
                }
                user = {
                    username: data[0],
                    sessionKey: data[1],
                    userSecret: data[2],
                }
                console.log("Loaded secrets successfully: ", user);
            } catch (e) {
                console.log("Error getting users secrets from db: ", e, data);
                return false
            }
            context.commit("SET_USERNAME", user.username, {root: true})
            context.commit("SET_SECRET_KEY", user.sessionKey, {root: true})
            context.commit("SET_USER_SECRET", user.userSecret, {root: true})
            return true;
        }
    },
    getters: {
        GET_SQLITE_CONN: state => {
            return state.conn;
        }
    }
}

export default sqliteDB;
