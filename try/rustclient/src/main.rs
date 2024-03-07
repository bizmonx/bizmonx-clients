use std::net::TcpStream;
use std::io::prelude::*;

fn main() -> std::io::Result<()> {
    // Connect to the server
    let mut stream = TcpStream::connect("localhost:1984")?;

    // Prepare the message
    let msg = "status newhost.rust green test from rust\n";

    // Send the message
    stream.write_all(msg.as_bytes())?;

    println!("Message sent successfully");
    Ok(())
}
