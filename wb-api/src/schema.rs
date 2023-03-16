use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct FilterOptions {
    pub page: Option<usize>,
    pub limit: Option<usize>,
}

#[derive(Deserialize, Debug)]
pub struct ParamOptions {
    pub id: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CreateCountrySchema {
    pub wbrn: String,
    #[serde(rename = "countryName")]
    pub country_name: String,   
}

#[derive(Serialize, Deserialize, Debug)]
pub struct UpdateCountrySchema{
    pub wbrn: String,
    #[serde(rename = "countryName")]
    pub country_name: String, 
}