use std::io::{BufRead, Write};

pub fn construct_uniqer(adjacent: bool, delete_lines: bool, ignore_case: bool) -> DefaultUniqer {
    let uniqer: Box<dyn Uniqer> = if adjacent {
        Box::new(AdjacentUniqer {
            prev: String::new(),
        })
    } else {
        Box::new(PlainUniqer { lines: vec![] })
    };
    return DefaultUniqer {
        uniqer: uniqer,
        ignore_case: ignore_case,
        delete_lines: delete_lines,
    };
}

pub trait Uniqer {
    fn next(&mut self, line: String) -> Option<String>;
}

pub struct DefaultUniqer {
    uniqer: Box<dyn Uniqer>,
    ignore_case: bool,
    delete_lines: bool,
}

impl DefaultUniqer {
    pub fn filter(
        &mut self,
        input: &mut Box<dyn BufRead>,
        output: &mut Box<dyn Write>,
    ) -> Result<(), String> {
        let mut buffer = String::new();
        loop {
            match input.read_line(&mut buffer) {
                Ok(0) => break,
                Ok(_) => {
                    match self.next(buffer.to_string()) {
                        None => (),
                        Some(line) => {
                            let _ = output.write(line.as_bytes());
                            ()
                        }
                    }
                    buffer.clear();
                }
                Err(why) => return Err(format!("{}", why)),
            };
        }
        Ok(())
    }
    fn line_or_none(&self, none_flag: bool, line: String) -> Option<String> {
        if none_flag {
            return None;
        }
        return Some(line);
    }
}

impl Uniqer for DefaultUniqer {
    fn next(&mut self, line: String) -> Option<String> {
        let target_line: String = if self.ignore_case {
            line.to_lowercase()
        } else {
            line.to_string()
        };
        let result = self.uniqer.next(target_line);
        return match result {
            Some(line) => self.line_or_none(self.delete_lines, line),
            None => self.line_or_none(!self.delete_lines, line),
        };
    }
}

struct AdjacentUniqer {
    prev: String,
}

impl Uniqer for AdjacentUniqer {
    fn next(&mut self, line: String) -> Option<String> {
        if self.prev == line {
            return None;
        }
        self.prev = line.to_string();
        return Some(line);
    }
}

struct PlainUniqer {
    lines: Vec<String>,
}

impl Uniqer for PlainUniqer {
    fn next(&mut self, line: String) -> Option<String> {
        if self.lines.contains(&line) {
            return None;
        }
        self.lines.push(line.to_string());
        return Some(line);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_plain_uniqer() {
        let mut uniqer = PlainUniqer {
            lines: vec![
                "line1".to_string(),
                "line2".to_string(),
                "line3".to_string(),
            ],
        };
        assert!(uniqer.next("line0".to_string()).is_some());
        assert!(uniqer.next("line1".to_string()).is_none());
        assert!(uniqer.next("line2".to_string()).is_none());
        assert!(uniqer.next("line3".to_string()).is_none());
    }

    #[test]
    fn test_adjacent_uniqer() {
        let mut uniqer = AdjacentUniqer {
            prev: String::new(),
        };
        assert!(uniqer.next("hogehoge".to_string()).is_some());
        assert!(uniqer.next("hogehoge".to_string()).is_none());
        assert!(uniqer.next("previous".to_string()).is_some());
        assert!(uniqer.next("previous".to_string()).is_none());
    }

    #[test]
    fn test_default_uniqer_with_adjacent_and_ignore_case() {
        let mut uniqer = construct_uniqer(true, false, true);
        assert!(uniqer.next("hogehoge".to_string()).is_some());
        assert!(uniqer.next("HOGEhoge".to_string()).is_none());
        assert!(uniqer.next("previous".to_string()).is_some());
        assert!(uniqer.next("previous".to_string()).is_none());
        assert!(uniqer.next("HOGEhoge".to_string()).is_some());
    }

    #[test]
    fn test_default_uniqer_with_none_adjacent_and_ignore_case() {
        let mut uniqer = construct_uniqer(false, false, true);
        assert!(uniqer.next("hogehoge".to_string()).is_some());
        assert!(uniqer.next("HOGEhoge".to_string()).is_none());
        assert!(uniqer.next("previous".to_string()).is_some());
        assert!(uniqer.next("previous".to_string()).is_none());
        assert!(uniqer.next("HOGEhoge".to_string()).is_none());
    }

    #[test]
    fn test_default_uniqer_with_adjacent_and_ignore_case_delete() {
        let mut uniqer = construct_uniqer(true, true, true);
        assert!(uniqer.next("hogehoge".to_string()).is_none());
        assert!(uniqer.next("HOGEhoge".to_string()).is_some());
        assert!(uniqer.next("previous".to_string()).is_none());
        assert!(uniqer.next("previous".to_string()).is_some());
        assert!(uniqer.next("HOGEhoge".to_string()).is_none());
    }

    #[test]
    fn test_default_uniqer_with_none_adjacent_and_ignore_case_delete() {
        let mut uniqer = construct_uniqer(false, true, true);
        assert!(uniqer.next("hogehoge".to_string()).is_none());
        assert!(uniqer.next("HOGEhoge".to_string()).is_some());
        assert!(uniqer.next("previous".to_string()).is_none());
        assert!(uniqer.next("previous".to_string()).is_some());
        assert!(uniqer.next("HOGEhoge".to_string()).is_some());
    }
}
