# Extract Text from JSON by GJSON

이 도구는 [`github.com/tidwall/gjson`](https://github.com/tidwall/gjson)을 이용해 JSON 파일에서 원하는 경로의 값을 추출한 뒤 텍스트 파일로 저장하는 간단한 CLI입니다. 파일 경로(`path`)와 GJSON 표현식(`rex`)만으로 빠르게 특정 필드 값을 뽑아낼 수 있으며, 선택적으로 사용 CPU 개수를 제어하는 `set_cpu` 인자를 제공합니다.

## 필수 및 선택 인자
- `path` (필수): 읽어올 JSON 파일의 경로입니다. Windows 경로 구분자(`\\`)는 자동으로 `/`로 치환되며, 따옴표(`"`)는 제거됩니다.
- `rex` (필수): [`gjson`](https://github.com/tidwall/gjson/blob/master/SYNTAX.md) 표현식입니다. 표현식이 배열을 반환하면 모든 요소가 출력 파일에 줄바꿈으로 나열됩니다.
- `set_cpu` (선택): `runtime.GOMAXPROCS` 값으로 설정할 CPU 개수입니다. 비워두면 1로 고정되며, 사용 가능한 CPU 수보다 큰 값을 주면 자동으로 최대치로 맞춰집니다.

## 동작 방식 요약
1. `path`로 지정된 파일을 읽어 전체 내용을 문자열로 로드합니다.
2. `rex` 표현식으로 값을 추출하고, 결과 배열을 문자열 슬라이스로 변환합니다.
3. 결과 문자열을 줄 단위로 합쳐 `<입력파일명>_<표현식 마지막 조각>.txt` 파일로 저장합니다.
4. 동일한 이름의 파일이 이미 존재하면 덮어쓰기를 묻습니다.

## 빠른 실행 예시
### 1. 샘플 JSON 준비
```json
{
  "posts": [
    {"title": "첫 번째 포스트", "tags": ["go", "json"]},
    {"title": "두 번째 포스트", "tags": ["cli"]}
  ]
}
```
위 내용을 `sample.json`으로 저장합니다.

### 2. 실행 명령
```bash
$ go run main.go sample.json "posts.#.title"
```

### 3. 생성되는 출력 파일
`sample_posts.#.title.txt` 파일이 생성되며 내용은 다음과 같습니다.
```
첫 번째 포스트
두 번째 포스트
```

## `main.go` 실행 시 출력 예시
```
$ go run main.go sample.json "posts.#.tags.#" 2
set cpu as 2
file exists, do you want to overwrite it? (y/n):
```
위 예시는 `set_cpu`를 2로 설정한 경우이며, 동일한 이름의 출력 파일이 이미 존재하면 덮어쓰기 여부를 묻습니다.

## 오류 처리 및 주의 사항
- `path`나 `rex`가 비어 있으면 프로그램은 즉시 종료하며 안내 메시지를 출력합니다.
- `rex` 결과가 비어 있으면 `string_array_to_file` 단계에서 패닉이 발생합니다. 표현식을 다시 확인하세요.
- 이미 존재하는 출력 파일은 덮어쓰기 전에 사용자 확인을 요청합니다. 비대화식 환경에서는 `y`를 제공하지 않으면 파일이 유지됩니다.
- 파일을 열거나 쓰는 과정에서 문제가 발생하면 표준 출력에 오류 메시지가 표시됩니다.
- `set_cpu` 인자에 숫자가 아닌 값이 들어오면 변환 오류가 출력되며 기본 CPU 설정이 유지됩니다.

## 개발 및 테스트 팁
- 모듈 의존성은 `go get` 없이도 `go run` 실행 시 자동으로 다운로드됩니다.
- 새 표현식을 테스트할 때는 `gjson` 프로젝트 문서의 예제를 참고하면 빠르게 작성할 수 있습니다.

## 라이선스
이 저장소의 라이선스 정책은 프로젝트 루트를 참고하세요.
