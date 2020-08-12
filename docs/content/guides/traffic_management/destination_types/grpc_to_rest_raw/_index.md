---
title: gRPC to REST Raw
weight: 135
description: Routing gRPC services to a REST API using raw Envoy config
---


```yaml
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  labels:
    app: gloo
  name: gateway-proxy
  namespace: gloo-system
spec:
  bindAddress: '::'
  bindPort: 8080
  httpGateway:
    options:
      grpc_json_transcoder:
        proto_descriptor_bin: Q3F3RkNoVm5iMjluYkdVdllYQnBMMmgwZEhBdWNISnZkRzhTQ21kdmIyZHNaUzVoY0draWVRb0VTSFIwY0JJcUNnVnlkV3hsY3hnQklBTW9DeklVTG1kdmIyZHNaUzVoY0drdVNIUjBjRkoxYkdWU0JYSjFiR1Z6RWtVS0gyWjFiR3g1WDJSbFkyOWtaVjl5WlhObGNuWmxaRjlsZUhCaGJuTnBiMjRZQWlBQktBaFNIR1oxYkd4NVJHVmpiMlJsVW1WelpYSjJaV1JGZUhCaGJuTnBiMjRpMmdJS0NFaDBkSEJTZFd4bEVob0tDSE5sYkdWamRHOXlHQUVnQVNnSlVnaHpaV3hsWTNSdmNoSVNDZ05uWlhRWUFpQUJLQWxJQUZJRFoyVjBFaElLQTNCMWRCZ0RJQUVvQ1VnQVVnTndkWFFTRkFvRWNHOXpkQmdFSUFFb0NVZ0FVZ1J3YjNOMEVoZ0tCbVJsYkdWMFpSZ0ZJQUVvQ1VnQVVnWmtaV3hsZEdVU0Znb0ZjR0YwWTJnWUJpQUJLQWxJQUZJRmNHRjBZMmdTTndvR1kzVnpkRzl0R0FnZ0FTZ0xNaDB1WjI5dloyeGxMbUZ3YVM1RGRYTjBiMjFJZEhSd1VHRjBkR1Z5YmtnQVVnWmpkWE4wYjIwU0Vnb0VZbTlrZVJnSElBRW9DVklFWW05a2VSSWpDZzF5WlhOd2IyNXpaVjlpYjJSNUdBd2dBU2dKVWd4eVpYTndiMjV6WlVKdlpIa1NSUW9UWVdSa2FYUnBiMjVoYkY5aWFXNWthVzVuY3hnTElBTW9DeklVTG1kdmIyZHNaUzVoY0drdVNIUjBjRkoxYkdWU0VtRmtaR2wwYVc5dVlXeENhVzVrYVc1bmMwSUpDZ2R3WVhSMFpYSnVJanNLRVVOMWMzUnZiVWgwZEhCUVlYUjBaWEp1RWhJS0JHdHBibVFZQVNBQktBbFNCR3RwYm1RU0Vnb0VjR0YwYUJnQ0lBRW9DVklFY0dGMGFFSnFDZzVqYjIwdVoyOXZaMnhsTG1Gd2FVSUpTSFIwY0ZCeWIzUnZVQUZhUVdkdmIyZHNaUzVuYjJ4aGJtY3ViM0puTDJkbGJuQnliM1J2TDJkdmIyZHNaV0Z3YVhNdllYQnBMMkZ1Ym05MFlYUnBiMjV6TzJGdWJtOTBZWFJwYjI1eitBRUJvZ0lFUjBGUVNXSUdjSEp2ZEc4ekNwczdDaUJuYjI5bmJHVXZjSEp2ZEc5aWRXWXZaR1Z6WTNKcGNIUnZjaTV3Y205MGJ4SVBaMjl2WjJ4bExuQnliM1J2WW5WbUlrMEtFVVpwYkdWRVpYTmpjbWx3ZEc5eVUyVjBFamdLQkdacGJHVVlBU0FES0FzeUpDNW5iMjluYkdVdWNISnZkRzlpZFdZdVJtbHNaVVJsYzJOeWFYQjBiM0pRY205MGIxSUVabWxzWlNMa0JBb1RSbWxzWlVSbGMyTnlhWEIwYjNKUWNtOTBieElTQ2dSdVlXMWxHQUVnQVNnSlVnUnVZVzFsRWhnS0IzQmhZMnRoWjJVWUFpQUJLQWxTQjNCaFkydGhaMlVTSGdvS1pHVndaVzVrWlc1amVSZ0RJQU1vQ1ZJS1pHVndaVzVrWlc1amVSSXJDaEZ3ZFdKc2FXTmZaR1Z3Wlc1a1pXNWplUmdLSUFNb0JWSVFjSFZpYkdsalJHVndaVzVrWlc1amVSSW5DZzkzWldGclgyUmxjR1Z1WkdWdVkza1lDeUFES0FWU0RuZGxZV3RFWlhCbGJtUmxibU41RWtNS0RHMWxjM05oWjJWZmRIbHdaUmdFSUFNb0N6SWdMbWR2YjJkc1pTNXdjbTkwYjJKMVppNUVaWE5qY21sd2RHOXlVSEp2ZEc5U0MyMWxjM05oWjJWVWVYQmxFa0VLQ1dWdWRXMWZkSGx3WlJnRklBTW9DeklrTG1kdmIyZHNaUzV3Y205MGIySjFaaTVGYm5WdFJHVnpZM0pwY0hSdmNsQnliM1J2VWdobGJuVnRWSGx3WlJKQkNnZHpaWEoyYVdObEdBWWdBeWdMTWljdVoyOXZaMnhsTG5CeWIzUnZZblZtTGxObGNuWnBZMlZFWlhOamNtbHdkRzl5VUhKdmRHOVNCM05sY25acFkyVVNRd29KWlhoMFpXNXphVzl1R0FjZ0F5Z0xNaVV1WjI5dloyeGxMbkJ5YjNSdlluVm1Ma1pwWld4a1JHVnpZM0pwY0hSdmNsQnliM1J2VWdsbGVIUmxibk5wYjI0U05nb0hiM0IwYVc5dWN4Z0lJQUVvQ3pJY0xtZHZiMmRzWlM1d2NtOTBiMkoxWmk1R2FXeGxUM0IwYVc5dWMxSUhiM0IwYVc5dWN4SkpDaEJ6YjNWeVkyVmZZMjlrWlY5cGJtWnZHQWtnQVNnTE1oOHVaMjl2WjJ4bExuQnliM1J2WW5WbUxsTnZkWEpqWlVOdlpHVkpibVp2VWc1emIzVnlZMlZEYjJSbFNXNW1ieElXQ2daemVXNTBZWGdZRENBQktBbFNCbk41Ym5SaGVDSzVCZ29QUkdWelkzSnBjSFJ2Y2xCeWIzUnZFaElLQkc1aGJXVVlBU0FCS0FsU0JHNWhiV1VTT3dvRlptbGxiR1FZQWlBREtBc3lKUzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVSbWxsYkdSRVpYTmpjbWx3ZEc5eVVISnZkRzlTQldacFpXeGtFa01LQ1dWNGRHVnVjMmx2YmhnR0lBTW9DeklsTG1kdmIyZHNaUzV3Y205MGIySjFaaTVHYVdWc1pFUmxjMk55YVhCMGIzSlFjbTkwYjFJSlpYaDBaVzV6YVc5dUVrRUtDMjVsYzNSbFpGOTBlWEJsR0FNZ0F5Z0xNaUF1WjI5dloyeGxMbkJ5YjNSdlluVm1Ma1JsYzJOeWFYQjBiM0pRY205MGIxSUtibVZ6ZEdWa1ZIbHdaUkpCQ2dsbGJuVnRYM1I1Y0dVWUJDQURLQXN5SkM1bmIyOW5iR1V1Y0hKdmRHOWlkV1l1Ulc1MWJVUmxjMk55YVhCMGIzSlFjbTkwYjFJSVpXNTFiVlI1Y0dVU1dBb1BaWGgwWlc1emFXOXVYM0poYm1kbEdBVWdBeWdMTWk4dVoyOXZaMnhsTG5CeWIzUnZZblZtTGtSbGMyTnlhWEIwYjNKUWNtOTBieTVGZUhSbGJuTnBiMjVTWVc1blpWSU9aWGgwWlc1emFXOXVVbUZ1WjJVU1JBb0tiMjVsYjJaZlpHVmpiQmdJSUFNb0N6SWxMbWR2YjJkc1pTNXdjbTkwYjJKMVppNVBibVZ2WmtSbGMyTnlhWEIwYjNKUWNtOTBiMUlKYjI1bGIyWkVaV05zRWprS0IyOXdkR2x2Ym5NWUJ5QUJLQXN5SHk1bmIyOW5iR1V1Y0hKdmRHOWlkV1l1VFdWemMyRm5aVTl3ZEdsdmJuTlNCMjl3ZEdsdmJuTVNWUW9PY21WelpYSjJaV1JmY21GdVoyVVlDU0FES0FzeUxpNW5iMjluYkdVdWNISnZkRzlpZFdZdVJHVnpZM0pwY0hSdmNsQnliM1J2TGxKbGMyVnlkbVZrVW1GdVoyVlNEWEpsYzJWeWRtVmtVbUZ1WjJVU0l3b05jbVZ6WlhKMlpXUmZibUZ0WlJnS0lBTW9DVklNY21WelpYSjJaV1JPWVcxbEdub0tEa1Y0ZEdWdWMybHZibEpoYm1kbEVoUUtCWE4wWVhKMEdBRWdBU2dGVWdWemRHRnlkQklRQ2dObGJtUVlBaUFCS0FWU0EyVnVaQkpBQ2dkdmNIUnBiMjV6R0FNZ0FTZ0xNaVl1WjI5dloyeGxMbkJ5YjNSdlluVm1Ma1Y0ZEdWdWMybHZibEpoYm1kbFQzQjBhVzl1YzFJSGIzQjBhVzl1Y3hvM0NnMVNaWE5sY25abFpGSmhibWRsRWhRS0JYTjBZWEowR0FFZ0FTZ0ZVZ1Z6ZEdGeWRCSVFDZ05sYm1RWUFpQUJLQVZTQTJWdVpDSjhDaFZGZUhSbGJuTnBiMjVTWVc1blpVOXdkR2x2Ym5NU1dBb1VkVzVwYm5SbGNuQnlaWFJsWkY5dmNIUnBiMjRZNXdjZ0F5Z0xNaVF1WjI5dloyeGxMbkJ5YjNSdlluVm1MbFZ1YVc1MFpYSndjbVYwWldSUGNIUnBiMjVTRTNWdWFXNTBaWEp3Y21WMFpXUlBjSFJwYjI0cUNRam9CeENBZ0lDQUFpS1lCZ29VUm1sbGJHUkVaWE5qY21sd2RHOXlVSEp2ZEc4U0Vnb0VibUZ0WlJnQklBRW9DVklFYm1GdFpSSVdDZ1p1ZFcxaVpYSVlBeUFCS0FWU0JtNTFiV0psY2hKQkNnVnNZV0psYkJnRUlBRW9EaklyTG1kdmIyZHNaUzV3Y205MGIySjFaaTVHYVdWc1pFUmxjMk55YVhCMGIzSlFjbTkwYnk1TVlXSmxiRklGYkdGaVpXd1NQZ29FZEhsd1pSZ0ZJQUVvRGpJcUxtZHZiMmRzWlM1d2NtOTBiMkoxWmk1R2FXVnNaRVJsYzJOeWFYQjBiM0pRY205MGJ5NVVlWEJsVWdSMGVYQmxFaHNLQ1hSNWNHVmZibUZ0WlJnR0lBRW9DVklJZEhsd1pVNWhiV1VTR2dvSVpYaDBaVzVrWldVWUFpQUJLQWxTQ0dWNGRHVnVaR1ZsRWlNS0RXUmxabUYxYkhSZmRtRnNkV1VZQnlBQktBbFNER1JsWm1GMWJIUldZV3gxWlJJZkNndHZibVZ2Wmw5cGJtUmxlQmdKSUFFb0JWSUtiMjVsYjJaSmJtUmxlQkliQ2dscWMyOXVYMjVoYldVWUNpQUJLQWxTQ0dwemIyNU9ZVzFsRWpjS0IyOXdkR2x2Ym5NWUNDQUJLQXN5SFM1bmIyOW5iR1V1Y0hKdmRHOWlkV1l1Um1sbGJHUlBjSFJwYjI1elVnZHZjSFJwYjI1eklyWUNDZ1JVZVhCbEVnOEtDMVJaVUVWZlJFOVZRa3hGRUFFU0Rnb0tWRmxRUlY5R1RFOUJWQkFDRWc0S0NsUlpVRVZmU1U1VU5qUVFBeElQQ2d0VVdWQkZYMVZKVGxRMk5CQUVFZzRLQ2xSWlVFVmZTVTVVTXpJUUJSSVFDZ3hVV1ZCRlgwWkpXRVZFTmpRUUJoSVFDZ3hVV1ZCRlgwWkpXRVZFTXpJUUJ4SU5DZ2xVV1ZCRlgwSlBUMHdRQ0JJUENndFVXVkJGWDFOVVVrbE9SeEFKRWc0S0NsUlpVRVZmUjFKUFZWQVFDaElRQ2d4VVdWQkZYMDFGVTFOQlIwVVFDeElPQ2dwVVdWQkZYMEpaVkVWVEVBd1NEd29MVkZsUVJWOVZTVTVVTXpJUURSSU5DZ2xVV1ZCRlgwVk9WVTBRRGhJUkNnMVVXVkJGWDFOR1NWaEZSRE15RUE4U0VRb05WRmxRUlY5VFJrbFlSVVEyTkJBUUVnOEtDMVJaVUVWZlUwbE9WRE15RUJFU0R3b0xWRmxRUlY5VFNVNVVOalFRRWlKRENnVk1ZV0psYkJJU0NnNU1RVUpGVEY5UFVGUkpUMDVCVEJBQkVoSUtEa3hCUWtWTVgxSkZVVlZKVWtWRUVBSVNFZ29PVEVGQ1JVeGZVa1ZRUlVGVVJVUVFBeUpqQ2hSUGJtVnZaa1JsYzJOeWFYQjBiM0pRY205MGJ4SVNDZ1J1WVcxbEdBRWdBU2dKVWdSdVlXMWxFamNLQjI5d2RHbHZibk1ZQWlBQktBc3lIUzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVUMjVsYjJaUGNIUnBiMjV6VWdkdmNIUnBiMjV6SXVNQ0NoTkZiblZ0UkdWelkzSnBjSFJ2Y2xCeWIzUnZFaElLQkc1aGJXVVlBU0FCS0FsU0JHNWhiV1VTUHdvRmRtRnNkV1VZQWlBREtBc3lLUzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVSVzUxYlZaaGJIVmxSR1Z6WTNKcGNIUnZjbEJ5YjNSdlVnVjJZV3gxWlJJMkNnZHZjSFJwYjI1ekdBTWdBU2dMTWh3dVoyOXZaMnhsTG5CeWIzUnZZblZtTGtWdWRXMVBjSFJwYjI1elVnZHZjSFJwYjI1ekVsMEtEbkpsYzJWeWRtVmtYM0poYm1kbEdBUWdBeWdMTWpZdVoyOXZaMnhsTG5CeWIzUnZZblZtTGtWdWRXMUVaWE5qY21sd2RHOXlVSEp2ZEc4dVJXNTFiVkpsYzJWeWRtVmtVbUZ1WjJWU0RYSmxjMlZ5ZG1Wa1VtRnVaMlVTSXdvTmNtVnpaWEoyWldSZmJtRnRaUmdGSUFNb0NWSU1jbVZ6WlhKMlpXUk9ZVzFsR2pzS0VVVnVkVzFTWlhObGNuWmxaRkpoYm1kbEVoUUtCWE4wWVhKMEdBRWdBU2dGVWdWemRHRnlkQklRQ2dObGJtUVlBaUFCS0FWU0EyVnVaQ0tEQVFvWVJXNTFiVlpoYkhWbFJHVnpZM0pwY0hSdmNsQnliM1J2RWhJS0JHNWhiV1VZQVNBQktBbFNCRzVoYldVU0Znb0diblZ0WW1WeUdBSWdBU2dGVWdadWRXMWlaWElTT3dvSGIzQjBhVzl1Y3hnRElBRW9DekloTG1kdmIyZHNaUzV3Y205MGIySjFaaTVGYm5WdFZtRnNkV1ZQY0hScGIyNXpVZ2R2Y0hScGIyNXpJcWNCQ2haVFpYSjJhV05sUkdWelkzSnBjSFJ2Y2xCeWIzUnZFaElLQkc1aGJXVVlBU0FCS0FsU0JHNWhiV1VTUGdvR2JXVjBhRzlrR0FJZ0F5Z0xNaVl1WjI5dloyeGxMbkJ5YjNSdlluVm1MazFsZEdodlpFUmxjMk55YVhCMGIzSlFjbTkwYjFJR2JXVjBhRzlrRWprS0IyOXdkR2x2Ym5NWUF5QUJLQXN5SHk1bmIyOW5iR1V1Y0hKdmRHOWlkV1l1VTJWeWRtbGpaVTl3ZEdsdmJuTlNCMjl3ZEdsdmJuTWlpUUlLRlUxbGRHaHZaRVJsYzJOeWFYQjBiM0pRY205MGJ4SVNDZ1J1WVcxbEdBRWdBU2dKVWdSdVlXMWxFaDBLQ21sdWNIVjBYM1I1Y0dVWUFpQUJLQWxTQ1dsdWNIVjBWSGx3WlJJZkNndHZkWFJ3ZFhSZmRIbHdaUmdESUFFb0NWSUtiM1YwY0hWMFZIbHdaUkk0Q2dkdmNIUnBiMjV6R0FRZ0FTZ0xNaDR1WjI5dloyeGxMbkJ5YjNSdlluVm1MazFsZEdodlpFOXdkR2x2Ym5OU0IyOXdkR2x2Ym5NU01Bb1FZMnhwWlc1MFgzTjBjbVZoYldsdVp4Z0ZJQUVvQ0RvRlptRnNjMlZTRDJOc2FXVnVkRk4wY21WaGJXbHVaeEl3Q2hCelpYSjJaWEpmYzNSeVpXRnRhVzVuR0FZZ0FTZ0lPZ1ZtWVd4elpWSVBjMlZ5ZG1WeVUzUnlaV0Z0YVc1bklwSUpDZ3RHYVd4bFQzQjBhVzl1Y3hJaENneHFZWFpoWDNCaFkydGhaMlVZQVNBQktBbFNDMnBoZG1GUVlXTnJZV2RsRWpBS0ZHcGhkbUZmYjNWMFpYSmZZMnhoYzNOdVlXMWxHQWdnQVNnSlVoSnFZWFpoVDNWMFpYSkRiR0Z6YzI1aGJXVVNOUW9UYW1GMllWOXRkV3gwYVhCc1pWOW1hV3hsY3hnS0lBRW9DRG9GWm1Gc2MyVlNFV3BoZG1GTmRXeDBhWEJzWlVacGJHVnpFa1FLSFdwaGRtRmZaMlZ1WlhKaGRHVmZaWEYxWVd4elgyRnVaRjlvWVhOb0dCUWdBU2dJUWdJWUFWSVphbUYyWVVkbGJtVnlZWFJsUlhGMVlXeHpRVzVrU0dGemFCSTZDaFpxWVhaaFgzTjBjbWx1WjE5amFHVmphMTkxZEdZNEdCc2dBU2dJT2dWbVlXeHpaVklUYW1GMllWTjBjbWx1WjBOb1pXTnJWWFJtT0JKVENneHZjSFJwYldsNlpWOW1iM0lZQ1NBQktBNHlLUzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVSbWxzWlU5d2RHbHZibk11VDNCMGFXMXBlbVZOYjJSbE9nVlRVRVZGUkZJTGIzQjBhVzFwZW1WR2IzSVNIUW9LWjI5ZmNHRmphMkZuWlJnTElBRW9DVklKWjI5UVlXTnJZV2RsRWpVS0UyTmpYMmRsYm1WeWFXTmZjMlZ5ZG1salpYTVlFQ0FCS0FnNkJXWmhiSE5sVWhGalkwZGxibVZ5YVdOVFpYSjJhV05sY3hJNUNoVnFZWFpoWDJkbGJtVnlhV05mYzJWeWRtbGpaWE1ZRVNBQktBZzZCV1poYkhObFVoTnFZWFpoUjJWdVpYSnBZMU5sY25acFkyVnpFalVLRTNCNVgyZGxibVZ5YVdOZmMyVnlkbWxqWlhNWUVpQUJLQWc2QldaaGJITmxVaEZ3ZVVkbGJtVnlhV05UWlhKMmFXTmxjeEkzQ2hSd2FIQmZaMlZ1WlhKcFkxOXpaWEoyYVdObGN4Z3FJQUVvQ0RvRlptRnNjMlZTRW5Cb2NFZGxibVZ5YVdOVFpYSjJhV05sY3hJbENncGtaWEJ5WldOaGRHVmtHQmNnQVNnSU9nVm1ZV3h6WlZJS1pHVndjbVZqWVhSbFpCSXZDaEJqWTE5bGJtRmliR1ZmWVhKbGJtRnpHQjhnQVNnSU9nVm1ZV3h6WlZJT1kyTkZibUZpYkdWQmNtVnVZWE1TS2dvUmIySnFZMTlqYkdGemMxOXdjbVZtYVhnWUpDQUJLQWxTRDI5aWFtTkRiR0Z6YzFCeVpXWnBlQklwQ2hCamMyaGhjbkJmYm1GdFpYTndZV05sR0NVZ0FTZ0pVZzlqYzJoaGNuQk9ZVzFsYzNCaFkyVVNJUW9NYzNkcFpuUmZjSEpsWm1sNEdDY2dBU2dKVWd0emQybG1kRkJ5WldacGVCSW9DaEJ3YUhCZlkyeGhjM05mY0hKbFptbDRHQ2dnQVNnSlVnNXdhSEJEYkdGemMxQnlaV1pwZUJJakNnMXdhSEJmYm1GdFpYTndZV05sR0NrZ0FTZ0pVZ3h3YUhCT1lXMWxjM0JoWTJVU05Bb1djR2h3WDIxbGRHRmtZWFJoWDI1aGJXVnpjR0ZqWlJnc0lBRW9DVklVY0dod1RXVjBZV1JoZEdGT1lXMWxjM0JoWTJVU0lRb01jblZpZVY5d1lXTnJZV2RsR0MwZ0FTZ0pVZ3R5ZFdKNVVHRmphMkZuWlJKWUNoUjFibWx1ZEdWeWNISmxkR1ZrWDI5d2RHbHZiaGpuQnlBREtBc3lKQzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVWVzVwYm5SbGNuQnlaWFJsWkU5d2RHbHZibElUZFc1cGJuUmxjbkJ5WlhSbFpFOXdkR2x2YmlJNkNneFBjSFJwYldsNlpVMXZaR1VTQ1FvRlUxQkZSVVFRQVJJTkNnbERUMFJGWDFOSldrVVFBaElRQ2d4TVNWUkZYMUpWVGxSSlRVVVFBeW9KQ09nSEVJQ0FnSUFDU2dRSUpoQW5JdEVDQ2c1TlpYTnpZV2RsVDNCMGFXOXVjeEk4Q2hkdFpYTnpZV2RsWDNObGRGOTNhWEpsWDJadmNtMWhkQmdCSUFFb0NEb0ZabUZzYzJWU0ZHMWxjM05oWjJWVFpYUlhhWEpsUm05eWJXRjBFa3dLSDI1dlgzTjBZVzVrWVhKa1gyUmxjMk55YVhCMGIzSmZZV05qWlhOemIzSVlBaUFCS0FnNkJXWmhiSE5sVWh4dWIxTjBZVzVrWVhKa1JHVnpZM0pwY0hSdmNrRmpZMlZ6YzI5eUVpVUtDbVJsY0hKbFkyRjBaV1FZQXlBQktBZzZCV1poYkhObFVncGtaWEJ5WldOaGRHVmtFaHNLQ1cxaGNGOWxiblJ5ZVJnSElBRW9DRklJYldGd1JXNTBjbmtTV0FvVWRXNXBiblJsY25CeVpYUmxaRjl2Y0hScGIyNFk1d2NnQXlnTE1pUXVaMjl2WjJ4bExuQnliM1J2WW5WbUxsVnVhVzUwWlhKd2NtVjBaV1JQY0hScGIyNVNFM1Z1YVc1MFpYSndjbVYwWldSUGNIUnBiMjRxQ1Fqb0J4Q0FnSUNBQWtvRUNBZ1FDVW9FQ0FrUUNpTGlBd29NUm1sbGJHUlBjSFJwYjI1ekVrRUtCV04wZVhCbEdBRWdBU2dPTWlNdVoyOXZaMnhsTG5CeWIzUnZZblZtTGtacFpXeGtUM0IwYVc5dWN5NURWSGx3WlRvR1UxUlNTVTVIVWdWamRIbHdaUklXQ2dad1lXTnJaV1FZQWlBQktBaFNCbkJoWTJ0bFpCSkhDZ1pxYzNSNWNHVVlCaUFCS0E0eUpDNW5iMjluYkdVdWNISnZkRzlpZFdZdVJtbGxiR1JQY0hScGIyNXpMa3BUVkhsd1pUb0pTbE5mVGs5U1RVRk1VZ1pxYzNSNWNHVVNHUW9FYkdGNmVSZ0ZJQUVvQ0RvRlptRnNjMlZTQkd4aGVua1NKUW9LWkdWd2NtVmpZWFJsWkJnRElBRW9DRG9GWm1Gc2MyVlNDbVJsY0hKbFkyRjBaV1FTR1FvRWQyVmhheGdLSUFFb0NEb0ZabUZzYzJWU0JIZGxZV3NTV0FvVWRXNXBiblJsY25CeVpYUmxaRjl2Y0hScGIyNFk1d2NnQXlnTE1pUXVaMjl2WjJ4bExuQnliM1J2WW5WbUxsVnVhVzUwWlhKd2NtVjBaV1JQY0hScGIyNVNFM1Z1YVc1MFpYSndjbVYwWldSUGNIUnBiMjRpTHdvRlExUjVjR1VTQ2dvR1UxUlNTVTVIRUFBU0NBb0VRMDlTUkJBQkVoQUtERk5VVWtsT1IxOVFTVVZEUlJBQ0lqVUtCa3BUVkhsd1pSSU5DZ2xLVTE5T1QxSk5RVXdRQUJJTkNnbEtVMTlUVkZKSlRrY1FBUklOQ2dsS1UxOU9WVTFDUlZJUUFpb0pDT2dIRUlDQWdJQUNTZ1FJQkJBRkluTUtERTl1Wlc5bVQzQjBhVzl1Y3hKWUNoUjFibWx1ZEdWeWNISmxkR1ZrWDI5d2RHbHZiaGpuQnlBREtBc3lKQzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVWVzVwYm5SbGNuQnlaWFJsWkU5d2RHbHZibElUZFc1cGJuUmxjbkJ5WlhSbFpFOXdkR2x2YmlvSkNPZ0hFSUNBZ0lBQ0lzQUJDZ3RGYm5WdFQzQjBhVzl1Y3hJZkNndGhiR3h2ZDE5aGJHbGhjeGdDSUFFb0NGSUtZV3hzYjNkQmJHbGhjeElsQ2dwa1pYQnlaV05oZEdWa0dBTWdBU2dJT2dWbVlXeHpaVklLWkdWd2NtVmpZWFJsWkJKWUNoUjFibWx1ZEdWeWNISmxkR1ZrWDI5d2RHbHZiaGpuQnlBREtBc3lKQzVuYjI5bmJHVXVjSEp2ZEc5aWRXWXVWVzVwYm5SbGNuQnlaWFJsWkU5d2RHbHZibElUZFc1cGJuUmxjbkJ5WlhSbFpFOXdkR2x2YmlvSkNPZ0hFSUNBZ0lBQ1NnUUlCUkFHSXA0QkNoQkZiblZ0Vm1Gc2RXVlBjSFJwYjI1ekVpVUtDbVJsY0hKbFkyRjBaV1FZQVNBQktBZzZCV1poYkhObFVncGtaWEJ5WldOaGRHVmtFbGdLRkhWdWFXNTBaWEp3Y21WMFpXUmZiM0IwYVc5dUdPY0hJQU1vQ3pJa0xtZHZiMmRzWlM1d2NtOTBiMkoxWmk1VmJtbHVkR1Z5Y0hKbGRHVmtUM0IwYVc5dVVoTjFibWx1ZEdWeWNISmxkR1ZrVDNCMGFXOXVLZ2tJNkFjUWdJQ0FnQUlpbkFFS0RsTmxjblpwWTJWUGNIUnBiMjV6RWlVS0NtUmxjSEpsWTJGMFpXUVlJU0FCS0FnNkJXWmhiSE5sVWdwa1pYQnlaV05oZEdWa0VsZ0tGSFZ1YVc1MFpYSndjbVYwWldSZmIzQjBhVzl1R09jSElBTW9DeklrTG1kdmIyZHNaUzV3Y205MGIySjFaaTVWYm1sdWRHVnljSEpsZEdWa1QzQjBhVzl1VWhOMWJtbHVkR1Z5Y0hKbGRHVmtUM0IwYVc5dUtna0k2QWNRZ0lDQWdBSWk0QUlLRFUxbGRHaHZaRTl3ZEdsdmJuTVNKUW9LWkdWd2NtVmpZWFJsWkJnaElBRW9DRG9GWm1Gc2MyVlNDbVJsY0hKbFkyRjBaV1FTY1FvUmFXUmxiWEJ2ZEdWdVkzbGZiR1YyWld3WUlpQUJLQTR5THk1bmIyOW5iR1V1Y0hKdmRHOWlkV1l1VFdWMGFHOWtUM0IwYVc5dWN5NUpaR1Z0Y0c5MFpXNWplVXhsZG1Wc09oTkpSRVZOVUU5VVJVNURXVjlWVGt0T1QxZE9VaEJwWkdWdGNHOTBaVzVqZVV4bGRtVnNFbGdLRkhWdWFXNTBaWEp3Y21WMFpXUmZiM0IwYVc5dUdPY0hJQU1vQ3pJa0xtZHZiMmRzWlM1d2NtOTBiMkoxWmk1VmJtbHVkR1Z5Y0hKbGRHVmtUM0IwYVc5dVVoTjFibWx1ZEdWeWNISmxkR1ZrVDNCMGFXOXVJbEFLRUVsa1pXMXdiM1JsYm1ONVRHVjJaV3dTRndvVFNVUkZUVkJQVkVWT1ExbGZWVTVMVGs5WFRoQUFFaE1LRDA1UFgxTkpSRVZmUlVaR1JVTlVVeEFCRWc0S0NrbEVSVTFRVDFSRlRsUVFBaW9KQ09nSEVJQ0FnSUFDSXBvRENoTlZibWx1ZEdWeWNISmxkR1ZrVDNCMGFXOXVFa0VLQkc1aGJXVVlBaUFES0FzeUxTNW5iMjluYkdVdWNISnZkRzlpZFdZdVZXNXBiblJsY25CeVpYUmxaRTl3ZEdsdmJpNU9ZVzFsVUdGeWRGSUVibUZ0WlJJcENoQnBaR1Z1ZEdsbWFXVnlYM1poYkhWbEdBTWdBU2dKVWc5cFpHVnVkR2xtYVdWeVZtRnNkV1VTTEFvU2NHOXphWFJwZG1WZmFXNTBYM1poYkhWbEdBUWdBU2dFVWhCd2IzTnBkR2wyWlVsdWRGWmhiSFZsRWl3S0VtNWxaMkYwYVhabFgybHVkRjkyWVd4MVpSZ0ZJQUVvQTFJUWJtVm5ZWFJwZG1WSmJuUldZV3gxWlJJaENneGtiM1ZpYkdWZmRtRnNkV1VZQmlBQktBRlNDMlJ2ZFdKc1pWWmhiSFZsRWlFS0RITjBjbWx1WjE5MllXeDFaUmdISUFFb0RGSUxjM1J5YVc1blZtRnNkV1VTSndvUFlXZG5jbVZuWVhSbFgzWmhiSFZsR0FnZ0FTZ0pVZzVoWjJkeVpXZGhkR1ZXWVd4MVpScEtDZ2hPWVcxbFVHRnlkQkliQ2dsdVlXMWxYM0JoY25RWUFTQUNLQWxTQ0c1aGJXVlFZWEowRWlFS0RHbHpYMlY0ZEdWdWMybHZiaGdDSUFJb0NGSUxhWE5GZUhSbGJuTnBiMjRpcHdJS0RsTnZkWEpqWlVOdlpHVkpibVp2RWtRS0NHeHZZMkYwYVc5dUdBRWdBeWdMTWlndVoyOXZaMnhsTG5CeWIzUnZZblZtTGxOdmRYSmpaVU52WkdWSmJtWnZMa3h2WTJGMGFXOXVVZ2hzYjJOaGRHbHZiaHJPQVFvSVRHOWpZWFJwYjI0U0Znb0VjR0YwYUJnQklBTW9CVUlDRUFGU0JIQmhkR2dTRmdvRWMzQmhiaGdDSUFNb0JVSUNFQUZTQkhOd1lXNFNLUW9RYkdWaFpHbHVaMTlqYjIxdFpXNTBjeGdESUFFb0NWSVBiR1ZoWkdsdVowTnZiVzFsYm5SekVpc0tFWFJ5WVdsc2FXNW5YMk52YlcxbGJuUnpHQVFnQVNnSlVoQjBjbUZwYkdsdVowTnZiVzFsYm5SekVqb0tHV3hsWVdScGJtZGZaR1YwWVdOb1pXUmZZMjl0YldWdWRITVlCaUFES0FsU0YyeGxZV1JwYm1kRVpYUmhZMmhsWkVOdmJXMWxiblJ6SXRFQkNoRkhaVzVsY21GMFpXUkRiMlJsU1c1bWJ4Sk5DZ3BoYm01dmRHRjBhVzl1R0FFZ0F5Z0xNaTB1WjI5dloyeGxMbkJ5YjNSdlluVm1Ma2RsYm1WeVlYUmxaRU52WkdWSmJtWnZMa0Z1Ym05MFlYUnBiMjVTQ21GdWJtOTBZWFJwYjI0YWJRb0tRVzV1YjNSaGRHbHZiaElXQ2dSd1lYUm9HQUVnQXlnRlFnSVFBVklFY0dGMGFCSWZDZ3R6YjNWeVkyVmZabWxzWlJnQ0lBRW9DVklLYzI5MWNtTmxSbWxzWlJJVUNnVmlaV2RwYmhnRElBRW9CVklGWW1WbmFXNFNFQW9EWlc1a0dBUWdBU2dGVWdObGJtUkNqd0VLRTJOdmJTNW5iMjluYkdVdWNISnZkRzlpZFdaQ0VFUmxjMk55YVhCMGIzSlFjbTkwYjNOSUFWbytaMmwwYUhWaUxtTnZiUzluYjJ4aGJtY3ZjSEp2ZEc5aWRXWXZjSEp2ZEc5akxXZGxiaTFuYnk5a1pYTmpjbWx3ZEc5eU8yUmxjMk55YVhCMGIzTDRBUUdpQWdOSFVFS3FBaHBIYjI5bmJHVXVVSEp2ZEc5aWRXWXVVbVZtYkdWamRHbHZiZ3FvQWdvY1oyOXZaMnhsTDJGd2FTOWhibTV2ZEdGMGFXOXVjeTV3Y205MGJ4SUtaMjl2WjJ4bExtRndhUm9WWjI5dloyeGxMMkZ3YVM5b2RIUndMbkJ5YjNSdkdpQm5iMjluYkdVdmNISnZkRzlpZFdZdlpHVnpZM0pwY0hSdmNpNXdjbTkwYnpwTENnUm9kSFJ3RWg0dVoyOXZaMnhsTG5CeWIzUnZZblZtTGsxbGRHaHZaRTl3ZEdsdmJuTVlzTXE4SWlBQktBc3lGQzVuYjI5bmJHVXVZWEJwTGtoMGRIQlNkV3hsVWdSb2RIUndRbTRLRG1OdmJTNW5iMjluYkdVdVlYQnBRaEJCYm01dmRHRjBhVzl1YzFCeWIzUnZVQUZhUVdkdmIyZHNaUzVuYjJ4aGJtY3ViM0puTDJkbGJuQnliM1J2TDJkdmIyZHNaV0Z3YVhNdllYQnBMMkZ1Ym05MFlYUnBiMjV6TzJGdWJtOTBZWFJwYjI1em9nSUVSMEZRU1dJR2NISnZkRzh6Q3NJSENoQndjbTkwYnk5bWFXeGxMbkJ5YjNSdkVoQnpiMnh2TG1WNFlXMXdiR1Z6TG5ZeEdoeG5iMjluYkdVdllYQnBMMkZ1Ym05MFlYUnBiMjV6TG5CeWIzUnZJbElLQkVsMFpXMFNFZ29FYm1GdFpSZ0JJQUVvQ1ZJRWJtRnRaUklnQ2d0a1pYTmpjbWx3ZEdsdmJoZ0NJQUVvQ1ZJTFpHVnpZM0pwY0hScGIyNFNGQW9GY0hKcFkyVVlBeUFCS0FGU0JYQnlhV05sSWowS0QwZGxkRWwwWlcxU1pYTndiMjV6WlJJcUNnUnBkR1Z0R0FFZ0FTZ0xNaFl1YzI5c2J5NWxlR0Z0Y0d4bGN5NTJNUzVKZEdWdFVnUnBkR1Z0SWlRS0RrZGxkRWwwWlcxU1pYRjFaWE4wRWhJS0JHNWhiV1VZQVNBQktBbFNCRzVoYldVaVFRb1JUR2x6ZEVsMFpXMXpVbVZ6Y0c5dWMyVVNMQW9GYVhSbGJYTVlBU0FES0FzeUZpNXpiMnh2TG1WNFlXMXdiR1Z6TG5ZeExrbDBaVzFTQldsMFpXMXpJaElLRUV4cGMzUkpkR1Z0YzFKbGNYVmxjM1FpUUFvU1EzSmxZWFJsU1hSbGJWSmxjM0J2Ym5ObEVpb0tCR2wwWlcwWUFTQUJLQXN5Rmk1emIyeHZMbVY0WVcxd2JHVnpMbll4TGtsMFpXMVNCR2wwWlcwaVB3b1JRM0psWVhSbFNYUmxiVkpsY1hWbGMzUVNLZ29FYVhSbGJSZ0JJQUVvQ3pJV0xuTnZiRzh1WlhoaGJYQnNaWE11ZGpFdVNYUmxiVklFYVhSbGJTSkFDaEpFWld4bGRHVkpkR1Z0VW1WemNHOXVjMlVTS2dvRWFYUmxiUmdCSUFFb0N6SVdMbk52Ykc4dVpYaGhiWEJzWlhNdWRqRXVTWFJsYlZJRWFYUmxiU0luQ2hGRVpXeGxkR1ZKZEdWdFVtVnhkV1Z6ZEJJU0NnUnVZVzFsR0FFZ0FTZ0pVZ1J1WVcxbE11NENDZ3hUZEc5eVpWTmxjblpwWTJVU1dRb0tRM0psWVhSbFNYUmxiUklqTG5OdmJHOHVaWGhoYlhCc1pYTXVkakV1UTNKbFlYUmxTWFJsYlZKbGNYVmxjM1FhSkM1emIyeHZMbVY0WVcxd2JHVnpMbll4TGtOeVpXRjBaVWwwWlcxU1pYTndiMjV6WlNJQUVsWUtDVXhwYzNSSmRHVnRjeElpTG5OdmJHOHVaWGhoYlhCc1pYTXVkakV1VEdsemRFbDBaVzF6VW1WeGRXVnpkQm9qTG5OdmJHOHVaWGhoYlhCc1pYTXVkakV1VEdsemRFbDBaVzF6VW1WemNHOXVjMlVpQUJKWkNncEVaV3hsZEdWSmRHVnRFaU11YzI5c2J5NWxlR0Z0Y0d4bGN5NTJNUzVFWld4bGRHVkpkR1Z0VW1WeGRXVnpkQm9rTG5OdmJHOHVaWGhoYlhCc1pYTXVkakV1UkdWc1pYUmxTWFJsYlZKbGMzQnZibk5sSWdBU1VBb0hSMlYwU1hSbGJSSWdMbk52Ykc4dVpYaGhiWEJzWlhNdWRqRXVSMlYwU1hSbGJWSmxjWFZsYzNRYUlTNXpiMnh2TG1WNFlXMXdiR1Z6TG5ZeExrZGxkRWwwWlcxU1pYTndiMjV6WlNJQVFnZGFCWEJ5YjNSdllnWndjbTkwYnpNPQ==
        services:
          - solo.examples.v1.StoreService
  proxyNames:
  - gateway-proxy
  ssl: false
  useProxyProto: false
```

```yaml
apiVersion: gateway.solo.io/v1
kind: VirtualService
metadata:
  name: default
  namespace: gloo-system
spec:
  virtualHost:
    routes:
    - matchers:
       - methods:
         - GET
         prefix: /solo.examples.v1.StoreService
      routeAction:
       single:
         upstream:
           name: default-grpcstore-demo-80
           namespace: gloo-system
```

A growing trend is to use gRPC internally as the communication protocol between micro-services. This has quite a few advantages. Some of those are:

1. Client and server stubs are auto generated
1. Efficient binary protocol (Google's protobufs)
1. Cross-language support as client and server libraries are available in many languages
1. HTTP based which plays well with existing firewalls and load balancers
1. Well supported with tooling around observability

While gRPC works great for internal micro-services, it may be desirable to have the internet facing API be a JSON\REST style API. This can happen for many reasons. among which are:

1. Keeping the API backwards compatible
1. Making the API more Web friendly
1. Supporting low-end devices such as IoT where gRPC is not supported.

Gloo allows you to define JSON/REST to your gRPC API so you can have the best of both worlds - outwards facing REST API and an internal gRPC API with no extra code.

With Gloo, there is no need to annotate your proto definitions with the `google.api.http` options. A simple gRPC proto will work.

---

## Overview

In this guide we will deploy a gRPC micro-service and transform its gRPC API to a REST API via Gloo.

Usually, to understand the details of the binary protobuf, a protobuf descriptor is needed. As this micro-service is built with server reflection enabled; together with Gloo's automatic function discovery functionality the required protobuf descriptor will be automatically discovered.

In this guide we are going to:

1. Deploy a gRPC demo service
1. Verify that the gRPC descriptors were indeed discovered
1. Add a Virtual Service creating a REST API that maps to the gRPC API
1. Verify that everything is working as expected

Let's get started!

### Prereqs

Install Gloo with Function Discovery Service (FDS) [blacklist mode]({{< versioned_link_path fromRoot="/installation/advanced_configuration/fds_mode/#configuring-the-fdsmode-setting" >}}) enabled

---

## Deploy the demo gRPC store

Create a deployment and a service:

```shell
kubectl create deployment grpcstore-demo --image=docker.io/soloio/grpcstore-demo
kubectl expose deployment grpcstore-demo --port 80 --target-port=8080
```

### Verify that gRPC functions were discovered
After a few seconds Gloo should have discovered the service with it's proto descriptor:

```shell
kubectl get upstream -n gloo-system default-grpcstore-demo-80 -o yaml
```

You should see output similar to this:

```yaml
apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  labels:
    app: grpcstore-demo
    discovered_by: kubernetesplugin
  name: default-grpcstore-demo-80
  namespace: gloo-system
spec:
  discoveryMetadata: {}
  kube:
    selector:
      app: grpcstore-demo
    serviceName: grpcstore-demo
    serviceNamespace: default
    servicePort: 80
    serviceSpec:
      grpc:
        descriptors: Q3F3RkNoVm5iMjluYkdVdllYQnBMMmgwZE…bTkwYnpNPQ==
        grpcServices:
        - functionNames:
          - CreateItem
          - ListItems
          - DeleteItem
          - GetItem
          packageName: solo.examples.v1
          serviceName: StoreService
status:
  reported_by: gloo
  state: 1

```

{{% notice note %}}
The descriptors field above was truncated for brevity.
{{% /notice %}}

As you can see Gloo's function discovery detected the gRPC functions on that service. 

### Create a REST to gRPC translation

Now we are ready to create the external REST to gRPC API. Please run the following command:

```shell
kubectl create -f - <<EOF
apiVersion: gateway.solo.io/v1
kind: VirtualService
metadata:
  name: default
  namespace: gloo-system
spec:
  virtualHost:
    routes:
    - matchers:
       - methods:
         - GET
         prefix: /items/
      routeAction:
       single:
         destinationSpec:
           grpc:
             function: GetItem
             package: solo.examples.v1
             parameters:
               path: /items/{name}
             service: StoreService
         upstream:
           name: default-grpcstore-demo-80
           namespace: gloo-system
    - matchers:
       - methods:
         - DELETE
         prefix: /items/
      routeAction:
       single:
         destinationSpec:
           grpc:
             function: DeleteItem
             package: solo.examples.v1
             parameters:
               path: /items/{name}
             service: StoreService
         upstream:
           name: default-grpcstore-demo-80
           namespace: gloo-system
    - matchers:
       - methods:
         - GET
         exact: /items
      routeAction:
       single:
         destinationSpec:
           grpc:
             function: ListItems
             package: solo.examples.v1
             service: StoreService
         upstream:
           name: default-grpcstore-demo-80
           namespace: gloo-system
    - matchers:
       - methods:
         - POST
         exact: /items
      routeAction:
       single:
         destinationSpec:
           grpc:
             function: CreateItem
             package: solo.examples.v1
             service: StoreService
         upstream:
           name: default-grpcstore-demo-80
           namespace: gloo-system
EOF
```

An explanation for the Virtual Service above:
We have defined four routes. Each route uses a {{< protobuf name="grpc.options.gloo.solo.io.DestinationSpec" display="gRPC destinationSpec" >}} to define REST routes to a gRPC service. When translating a REST API to a gRPC API the JSON body is automatically used to fill in the proto message fields. If you have some parameters in the path or in headers, your can specify them using the {{< protobuf name="transformation.options.gloo.solo.io.Parameters" display="parameters">}}  block in the {{< protobuf name="grpc.options.gloo.solo.io.DestinationSpec" display="gRPC destinationSpec">}} (as done in the route to `GetItem` and `DeleteItem`). We use HTTP method matching to make sure that our API adheres to the REST semantics. Note that the routes for `CreateItem` and `ListItems` are defined for the exact path `/items` (i.e. no trailing slash).

### Test

To test, we can use `curl` to issue queries to our new REST API:

```shell
URL=$(glooctl proxy url)
# Create an item in the store.
curl $URL/items -d '{"item":{"name":"item1"}}'
# List all items in the store. You should see an object with a list containing the item created above. 
curl $URL/items
# Access a specific item. You should see the item as a single object.
curl $URL/items/item1
# Delete the item created.
curl $URL/items/item1 -XDELETE
# No items - this will return an empty object.
curl $URL/items
```

---

## Conclusion

In this guide we have deployed a gRPC micro-service and created an external REST API that translates to the gRPC API via Gloo. This allows you to enjoy the benefits of using gRPC for your microservices while still having a traditional REST API without the need to maintain two sets of code. 

### Next Steps

Learn more about how Gloo handles [gRPC for web clients]({{% versioned_link_path fromRoot="/guides/traffic_management/listener_configuration/grpc_web/" %}}). Gloo can also use a [REST endpoint]({{% versioned_link_path fromRoot="/guides/traffic_management/destination_types/rest_endpoint/" %}}) as an Upstream. Our [function discovery guide]({{% versioned_link_path fromRoot="/installation/advanced_configuration/fds_mode/" %}}) covers how to set up the Function Discovery Service (FDS) for a Swagger document or gRPC service.
