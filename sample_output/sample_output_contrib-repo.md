| SL NO |                                                          PR                                                          |             TITLE              |
|-------|----------------------------------------------------------------------------------------------------------------------|--------------------------------|
| 1.    | [prometheus-operator/prometheus-operator#4395](https://github.com/prometheus-operator/prometheus-operator/pull/4395) | pkg/prometheus/promcfg*.go:    |
|       |                                                                                                                      | Fix regex in relabel_configs   |
| 2.    | [prometheus-operator/prometheus-operator#4352](https://github.com/prometheus-operator/prometheus-operator/pull/4352) | Add validations for duration   |
|       |                                                                                                                      | fields in alertmanager and     |
|       |                                                                                                                      | thanos                         |
| 3.    | [prometheus-operator/prometheus-operator#4350](https://github.com/prometheus-operator/prometheus-operator/pull/4350) | pkg: Remove app label from     |
|       |                                                                                                                      | statefulsets/pods              |
| 4.    | [prometheus-operator/prometheus-operator#4327](https://github.com/prometheus-operator/prometheus-operator/pull/4327) | pkg/prometheus/promcfg*.go:    |
|       |                                                                                                                      | Allow matchLabels selector to  |
|       |                                                                                                                      | have empty label values        |
| 5.    | [prometheus-operator/prometheus-operator#4314](https://github.com/prometheus-operator/prometheus-operator/pull/4314) | Add .gitattributes             |
| 6.    | [prometheus-operator/prometheus-operator#4313](https://github.com/prometheus-operator/prometheus-operator/pull/4313) | Update to go 1.17 and format   |
| 7.    | [prometheus-operator/prometheus-operator#4308](https://github.com/prometheus-operator/prometheus-operator/pull/4308) | pkg/prometheus/promcfg.go: Add |
|       |                                                                                                                      | validation for duration and    |
|       |                                                                                                                      | size fields                    |
| 8.    | [prometheus-operator/prometheus-operator#4285](https://github.com/prometheus-operator/prometheus-operator/pull/4285) | pkg/prometheus/promcfg*.go:    |
|       |                                                                                                                      | Validate EnforcedBodySizeLimit |
| 9.    | [prometheus-operator/prometheus-operator#4275](https://github.com/prometheus-operator/prometheus-operator/pull/4275) | Support enforced               |
|       |                                                                                                                      | body_size_limit at Prometheus  |
|       |                                                                                                                      | CRD Level                      |
| 10.   | [prometheus-operator/prometheus-operator#4245](https://github.com/prometheus-operator/prometheus-operator/pull/4245) | pkg/prometheus/promcfg.go:     |
|       |                                                                                                                      | Remove check for prometheus    |
|       |                                                                                                                      | 1.x                            |
| 11.   | [prometheus-operator/prometheus-operator#4207](https://github.com/prometheus-operator/prometheus-operator/pull/4207) | Add sample and target limits   |
|       |                                                                                                                      | to Probe                       |
| 12.   | [prometheus-operator/prometheus-operator#4195](https://github.com/prometheus-operator/prometheus-operator/pull/4195) | Support label_limit,           |
|       |                                                                                                                      | label_name_length_limit and    |
|       |                                                                                                                      | label_value_length_limit       |
| 13.   | [prometheus-operator/prometheus-operator#4183](https://github.com/prometheus-operator/prometheus-operator/pull/4183) | Remove context field from      |
|       |                                                                                                                      | framework struct in tests      |
| 14.   | [prometheus-operator/prometheus-operator#4117](https://github.com/prometheus-operator/prometheus-operator/pull/4117) | Remove context.TODO() from     |
|       |                                                                                                                      | client_go function calls       |
| 15.   | [prometheus-operator/prometheus-operator#4104](https://github.com/prometheus-operator/prometheus-operator/pull/4104) | Remove context.TODO() from     |
|       |                                                                                                                      | client_go function calls       |
| 16.   | [prometheus-operator/prometheus-operator#4101](https://github.com/prometheus-operator/prometheus-operator/pull/4101) | pkg/prometheus/operator.go:    |
|       |                                                                                                                      | Update logger message          |
| 17.   | [prometheus-operator/prometheus-operator#4100](https://github.com/prometheus-operator/prometheus-operator/pull/4100) | operator.go: Update failure    |
|       |                                                                                                                      | message shown to user          |
