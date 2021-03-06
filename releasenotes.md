- [v0.12.0](#v0120)
- [v0.11.0](#v0110)
- [v0.10.0](#v0100)
- [v0.9.0](#v090)
- [v0.8.0](#v080)
- [v0.7.0](#v070)
- [v0.6.1](#v061)
- [v0.6.0](#v060)
- [v0.5.0](#v050)
- [v0.4.0](#v040)
- [v0.3.0](#v030)
- [v0.2.1](#v021)
- [v0.2.0](#v020)

## v0.12.0

### Enhancement
* Add Kubernetes 1.18 #139

### Bugfixes
* Fix a race condition when downloading pupernetes #142


## v0.11.0

### Enhancement
* Add Kubernetes 1.15 - 1.17, default is now 1.16.3 #137


## v0.10.0

### Enhancement
* kube-apiserver: remove deprecated admission control flag #135
* setup: increase the verbosity of everything #133
* Update CNI #132
* Add Kubernetes 1.14 #130


### Bugfixes
* Fix systemd unit files #131


## v0.9.0

### Enhancement
* node: support shared PID #128
* add kubernetes 1.13 #126
* setup: can choose the vault listen address #125
* version: introduce the display of it #121
* containerd: can choose the version over the command line #122
* binaries: can skip the verification of the version #120
* Support Fedora (systemd cgroupfs) #118
* versions: upgrade hyperkube, etcd and containerd #116

### Bugfixes
* resolv: remove the extension .conf #117
* containerd: kill move to process #115


### Other
* bump travis matrix #127
* makefile: set the constrain to go 1.10+ #123
* Update the license file #119
* examples: upgrade to the latest pupernetes #114


## v0.8.0

### Enhancement
* Create Kubernetes 1.12 template collection #107
* Add the admission control config file #109
* Introduce the keep verb #108
* Extend the audit logs to everything #110
* Remove useless stage in audit logs #106

### Bugfixes

### Other
* Refactor options package to use sets. #111

## v0.7.0

### Enhancement
* Skip useless binaries downloads #104
* Remove default timeout #102
* Support for containerd, use cni everywhere #90
* Introduce skip probe #100
* API aggregation, remove intermediate CA #98
* Introduce the audit logs to dir #99
* Use kube proxy configuration file #93
* Introduce a timeout and a sig handler during downloads #91
* Logging improvements on setup requirements #89

### Bugfixes
* Do create the directory for pod logs #96

### Other
* Fix typo #101
* Fix a typo in the dig command #97
* Improve circleci #92
* Refactor CI, introduce examples #71
* Fix ignition path #83

## v0.6.1

### Bugfixes
* Fix the re-apply system #84

## v0.6.0

### Enhancement
* Add a way to check if the dns is ready (#81)
* Stop the systemd units in a reverse order (#76)
* Remove the archive on extract failure (#74)
* Enable HPA (#68)

### Bugfixes
* job: fix the abspath of pupernetes, propagate SUDO_USER to the unit environment (#79)
* manifests: remove the duplicated flag (#75)

### Other
* Introduce release notes (#80)
* Fixed README and added new subsection for DNS requirement (#78)
* ci: add hyperkube 1.11 (#77)
* Introduce sonobuoy in CI (#69)

## v0.5.0

### Enhancement
* Add pprof #62
* Use RSA standard library #60
* Basic check on the memory #59
* Add a prometheus exporter #57
* Run over old systemd platform #49
* Command line daemon alias #48
* Add notify to the systemd unit #47 

### Bugfixes
* Use a dedicated timeout for each package #64
* Fix the signal reset on SIGs during Stop #58 

### Other
* Update the readme #63
* Introduce ineffassign, golint and misspell #56
* Add SaaS CI examples #50

## v0.4.0

### Enhancement
* New wait command #43

### Bugfixes
* Use a more portable version of listUnits #44

## v0.3.0

### Enhancement
* Introduce the daemon and the reset commands #41
* Allow to delete jobs #38 
* Configurable Kubernetes version setup #37 
* Change the severity of the notify when not running in systemd unit #34 
* Display more state of the runtime #33 

### Bugfixes
* Use an intermediate Certificate Authority for cluster-signing #36 
* Display the adapted config in the logs #35 
* Add aws public and ipv4 detection logic #29 

## v0.2.1

### Bugfixes
* Add aws public and ipv4 detection logic #29

## v0.2.0

### Enhancement
* Job type systemd: display the logs #26

### Bugfixes
* Fallback to AWS hostname metadata #23

### Other
* Disable the reboot strategy #22
* Refactor the Makefile #25
* Provide a release documentation #21
