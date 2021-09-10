use clap::Clap;

mod cli;
mod io;
mod uniqer;

#[cfg(not(tarpaulin_include))]
fn main() {
    let opts = cli::Opts::parse();

    let mut uniqer = uniqer::construct_uniqer(opts.adjacent, opts.delete_lines, opts.ignore_case);
    let r = match io::open(opts.input, opts.output) {
        Ok(tupple) => {
            let mut input = tupple.0;
            let mut output = tupple.1;
            uniqer.filter(&mut input, &mut output)
        }
        Err(e) => Err(format!("{}", e)),
    };
    match r {
        Err(e) => println!("{}", e),
        Ok(_) => (),
    };
}
