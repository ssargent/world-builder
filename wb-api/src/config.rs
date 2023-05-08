use crate::errors::Error;
use serde::{Deserialize, Serialize};
use url::Url;

const ENV_DATABASE_URL: &str = "DATABASE_URL";
const ENV_DATABASE_POOL_SIZE: &str = "DATABASE_POOL_SIZE";
/*const ENV_HTTP_PORT: &str = "PORT";
const ENV_HTTP_ACCESS_LOGS: &str = "HTTP_ACCESS_LOGS";
const ENV_HTTP_PUBLIC_DIRECTORY: &str = "HTTP_PUBLIC_DIRECTORY";
*/
const POSTGRES_SCHEME: &str = "postgres";

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Config {
    //    pub http: Http,
    pub database: Database,
}

/// Database contains the data necessary to connect to a database
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Database {
    pub url: String,
    pub pool_size: u32,
}
const DEFAULT_DATABASE_POOL_SIZE: u32 = 100;

impl Config {
    /// Load and validate the configuration from the environment.
    /// If an error is found while parsing the values, or validating the data, an error is returned.
    pub fn load() -> Result<Config, Error> {
        dotenv::dotenv().ok();

        // http
        /*        let http_port = std::env::var(ENV_HTTP_PORT)
                   .ok()
                   .map_or(Ok(DEFAULT_HTTP_PORT), |env_val| env_val.parse::<u16>())?;
               let http_access_logs = std::env::var(ENV_HTTP_ACCESS_LOGS)
                   .ok()
                   .map_or(Ok(DEFAULT_ACCESS_LOGS), |env_val| env_val.parse::<bool>())?;
               let http_public_directory =
                   std::env::var(ENV_HTTP_PUBLIC_DIRECTORY).unwrap_or(String::from(DEFAULT_HTTP_PUBLIC_DIRECTORY));

               let http = Http {
                   port: http_port,
                   access_logs: http_access_logs,
                   public_directory: http_public_directory,
               };
        */
        // database
        let database_url =
            std::env::var(ENV_DATABASE_URL).map_err(|_| env_not_found(ENV_DATABASE_URL))?;
        let database_pool_size = std::env::var(ENV_DATABASE_POOL_SIZE)
            .ok()
            .map_or(Ok(DEFAULT_DATABASE_POOL_SIZE), |pool_size_str| {
                pool_size_str.parse::<u32>()
            })?;

        let database = Database {
            url: database_url,
            pool_size: database_pool_size,
        };

        let mut config = Config {
            //          http,
            database,
        };

        config.clean_and_validate()?;

        Ok(config)
    }

    fn clean_and_validate(&mut self) -> Result<(), Error> {
        // Database
        let database_url = Url::parse(&self.database.url)?;
        if database_url.scheme() != POSTGRES_SCHEME {
            return Err(Error::InvalidArgument(String::from(
                "config: database_url is not a valid postgres URL",
            )));
        }

        Ok(())
    }
}

fn env_not_found(var: &str) -> Error {
    Error::NotFound(format!("config: {} env var not found", var))
}

#[cfg(test)]
pub mod test {
    use super::Config;
    use std::env;

    pub fn load_test_config() -> Config {
        // by default cargo run the tests in the directory of the package's manifest
        // in ou case, we need to run the tests in the directory of the workspace, in
        // order to load assets
        let current_dir = env::current_dir().unwrap();
        env::set_current_dir(current_dir.join("..")).unwrap();
        Config::load().unwrap()
    }
}
