use quicli::prelude::*;
use structopt::StructOpt;

/// run is a "meta" command-line tool -- its sole purpose is to run other
/// command-line tools in useful, interactive ways.
#[derive(Debug, StructOpt)]
struct Cli {
    // Add a CLI argument `--count`/-n` that defaults to 3, and has this help text:
    /// How many lines to get
    #[structopt(long = "count", short = "n", default_value = "3")]
    count: usize,

    /// A user-created route file for a given CLI tool.
    #[structopt(parse(from_os_str))]
    route: std::path::PathBuf,

    // Quick and easy logging setup you get for free with quicli
    #[structopt(flatten)]
    verbosity: Verbosity,
}

fn main() -> CliResult {
    let args = Cli::from_args();
    args.verbosity.setup_env_logger("run")?;

    let content = read_file(&args.route)?;
    let content_lines = content.lines();
    let first_n_lines = content_lines.take(args.count);

    info!("Reading first {} lines of {:?}", args.count, args.route);

    for line in first_n_lines {
        println!("{}", line);
    }

    Ok(())
}