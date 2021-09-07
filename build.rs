// in build.rs
use clap::IntoApp;
use clap_generate::{generate_to, generators::*};

include!("src/cli.rs");

fn main() {
    let mut app = Opts::into_app();
    app.set_bin_name("uniq2");

    let outdir = std::path::Path::new(env!("CARGO_MANIFEST_DIR")).join("completions");
    println!("outdir: {}", &outdir.display());
    let _ = generate_to::<Bash, _, _>(&mut app, "uniq2", &outdir);
    let _ = generate_to::<Fish, _, _>(&mut app, "uniq2", &outdir);
    let _ = generate_to::<Zsh, _, _>(&mut app, "uniq2", &outdir);
    let _ = generate_to::<PowerShell, _, _>(&mut app, "uniq2", &outdir);
    let _ = generate_to::<Elvish, _, _>(&mut app, "uniq2", &outdir);
}
