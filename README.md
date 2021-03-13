# go-quera.ir-jobs-scraper

> یک برنامه ی ساده برای خواندن همه ی شغل های موجود بر روی صفحه ی استخدام سایت
> [کوئرا](https://quera.ir/careers/jobs)
> می باشد

## Requirements

> دستور زیر را در
> ترمینال
> تایپ کرده و
> اینتر
> را بزنید تا بسته مورد نیاز نصب شود

- `go get -u github.com/gocolly/colly/...`

### Data Types

```go
// Job a data struct for select data
type Job struct {
	Title   string `json:"title"`
	Company string `json:"company"`
	Collab  string `json:"collab"`
	Salary  string `json:"salary"`
	Remote  string `json:"remote"`
}
```

And **JSON** output in `data(...).json`:

```json
{
  "title": "Senior Android Developer (Java)ارشد (Senior)",
  "company": "انتشارات آموزشی آرامش - تهران",
  "collab": "تمام‌وقت",
  "salary": "",
  "remote": ""
}
```
