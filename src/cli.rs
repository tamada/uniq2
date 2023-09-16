use clap::{Clap, ValueHint};

#[derive(Clap, Debug)]
#[clap(author, version, about, setting(clap::AppSettings::ColoredHelp))]
pub struct Opts {
    #[clap(short, long, about = "Delete only adjacent duplicated lines.")]
    pub adjacent: bool,
    #[clap(short, long, about = "Only prints deleted lines.")]
    pub delete_lines: bool,
    #[clap(short, long, about = "Case insensitive.")]
    pub ignore_case: bool,
    #[clap(
        name = "INPUT",
        value_hint = ValueHint::FilePath,
        about = "gives file name of input.  If argument is single dash ('-') or absent, the program read strings from stdin."
    )]
    pub input: Option<String>,
    #[clap(
        name = "OUTPUT",
        value_hint = ValueHint::FilePath,
        about = "represents the destination.  To specify this argument, must specify INPUT argument."
    )]
    pub output: Option<String>,
}
