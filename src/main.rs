#![allow(non_snake_case)]

use exitfailure::ExitFailure;
use serde_derive::{Deserialize, Serialize};
//use structopt::StructOpt;
use reqwest::Url;
use std::env;

#[derive(Serialize, Deserialize, Debug)]
struct FetchUser {
    userId: i32,
    id: i32,
    title: String,
    body: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct Backend {
    fetchUser: FetchUser,
}

impl Backend {
    async fn fetchUser(userId: i32) -> Result<FetchUser, ExitFailure> {
        let url: String = format!("https://jsonplaceholder.typicode.com/posts/{}", userId);
        let Url = Url::parse(&*url)?;

        let resp = reqwest::get(Url).await?.json::<FetchUser>().await?;
        return Ok(resp);
    }
}

// api.openweathermap.org/data/2.5/find?q=London&appid=78b1a361a6775e0485bbca6280b586d2
#[tokio::main]
async fn main() -> Result<(), ExitFailure> {
    let args: Vec<String> = env::args().collect();
    let response =
        Backend::fetchUser(args[1].parse().expect("failed to farse str into integer")).await?;
    println!("Response {:#?}", response);

    return Ok(());
}
