# Alfred LaTeX Table

Generate LaTeX Table with Alfred Workflow.

![screenshot](/screenshot.png)

## Installation

### Install manually

Download and import `alfred-latex-table.alfredworkflow`.

## Dependencies

- [Alfred](https://www.alfredapp.com/)
- [Alfred Powerpack](https://www.alfredapp.com/powerpack/)

## Usage

Generate empty LaTeX table:

```shell
Latex Table 3 3

Output:
\begin{table}[htbp]
\centering
\begin{tabular}{|c|c|c|}
\hline
  &   &   \\
\hline
  &   &   \\
  &   &   \\
  &   &   \\
\hline
\end{tabular}
\caption{}
\label{tab:}
\end{table}
```

Generate LaTeX table with test data:

```shell
Latex Table 3 3 test
# `Latex Table 3 3 data` is an alternative

Output:
\begin{table}[htbp]
\centering
\begin{tabular}{|c|c|c|}
\hline
COMPASSIONATE ALBATTANI & INSPIRING RITCHIE & UPBEAT BABBAGE \\
\hline
determined_haibt & goofy_darwin & hungry_wilson \\
elegant_rosalind & suspicious_swirles & flamboyant_dijkstra \\
compassionate_blackwell & elastic_franklin & silly_bardeen \\
\hline
\end{tabular}
\caption{}
\label{tab:}
\end{table}
```

## Credit

- [deanishe/awgo](https://github.com/deanishe/awgo)

## License

MIT
