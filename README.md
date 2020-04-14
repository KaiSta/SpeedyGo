# SpeedyGo

SpeedyGo is a framework to develop and evaluate race prediction algorithms. It includes a comprehensive benchmark to compare different algorithms in terms of false positives and false negatives.

## Usage

To run one of the benchmark tests with the W3PO algorithm, we use the following command:

```
go run main.go -mode w3po -trace ../benchmark/Test1.log
```

To get a list of supported algorithms:

```
go run main.go -ls
```

## List of Algorithms

* [SHBEE](https://arxiv.org/abs/1909.03289)
* [W3PO](https://github.com/KaiSta/SpeedyGo)
* [SSHB](https://arxiv.org/abs/1905.10855)
* SHB
* FastTrack
* WCP
* ThreadSanitizer
