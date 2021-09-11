use std::fs::{File, OpenOptions};
use std::io::{stdin, stdout, BufRead, BufReader, BufWriter, Write};
use std::path::Path;

fn open_input_file(file_name: &str) -> Result<Box<dyn BufRead>, String> {
    if file_name.trim() == "-" {
        return open_stdin();
    }
    let input_path = Path::new(file_name);
    if !input_path.exists() {
        return Err(format!("{}: file not found", &input_path.display()));
    }
    match File::open(&input_path) {
        Ok(file) => Ok(Box::new(BufReader::new(file))),
        Err(e) => Err(format!("{}: {}", input_path.display(), e.to_string())),
    }
}

fn open_stdin() -> Result<Box<dyn BufRead>, String> {
    return Ok(Box::new(BufReader::new(stdin())));
}

fn open_input(input: Option<String>) -> Result<Box<dyn BufRead>, String> {
    return if let Some(input) = input {
        open_input_file(&input)
    } else {
        open_stdin()
    };
}

fn open_output_file(output: &str) -> Result<Box<dyn Write>, String> {
    match OpenOptions::new().write(true).create(true).open(output) {
        Ok(file) => Ok(Box::new(BufWriter::new(file))),
        Err(e) => Err(format!("{}: {}", output, e.to_string())),
    }
}

fn open_stdout() -> Result<Box<dyn Write>, String> {
    Ok(Box::new(BufWriter::new(stdout())))
}

fn open_output(output: Option<String>) -> Result<Box<dyn Write>, String> {
    if let Some(output) = output {
        open_output_file(&output)
    } else {
        open_stdout()
    }
}

pub fn open(
    input: Option<String>,
    output: Option<String>,
) -> Result<(Box<dyn BufRead>, Box<dyn Write>), String> {
    let input = open_input(input)?;
    let output = open_output(output)?;
    Ok((input, output))
}

mod tests {
    use super::*;

    #[test]
    fn test_input_with_none() {
        let input = open_input(None);
        assert!(input.is_ok());
    }

    #[test]
    fn test_input_with_minus() {
        let input = open_input(Some("-".to_string()));
        assert!(input.is_ok());
    }

    #[test]
    fn test_input_with_file() {
        let input = open_input(Some("testdata/test1.txt".to_string()));
        assert!(input.is_ok());
    }

    #[test]
    fn test_input_with_no_exist_file() {
        let input = open_input(Some("testdata/not_exist_file.txt".to_string()));
        assert!(input.is_err());
    }

    #[test]
    fn test_output_with_not_exist_file() {
        let output = open_output(Some("testdata/not_exist_file.txt".to_string()));
        assert!(output.is_ok());
        let _ = std::fs::remove_file("testdata/not_exist_file.txt");
    }

    #[test]
    fn test_output_stdin() {
        let output = open_output(None);
        assert!(output.is_ok());
    }
}
