# 💪 BMI Calculator

## 💻 Requirements:

### [`Docker`](https://www.docker.com/) & [`Docker compose`](https://docs.docker.com/compose/)

## Installation & Running:

### Docker

```bash
docker compose up -d
```

## 💅 UI

[BMI Calculator App UI](https://www.figma.com/design/YOC2GryjTBeKukeqL6agTN/BMI-Calculator-App-UI--Community-?node-id=0-1&p=f&t=K832e0kZmfI4lv7k-0)

## 🎞️ Preview

<details>
    <summary>🖥️ Desktop</summary>
    <img alt="bmi calculator desktop" src="https://imgur.com/WayK09v.gif">
</details>

<details>
    <summary>📱 Mobile</summary>
    <img alt="bmi calculator mobile" src="https://imgur.com/ZZiN1tM.gif">
</details>


## 📜 Documentation:

|Service|Host Port|Container Port|Description|Observations|
|:---:|:---:|:---:|:---:|:---:|
|frontend|3000|3000|Next.js application|
|backend|8000|8000|Go Api|if BMI_API_DEV_PORT is set in `host` documentation is available in `/docs` endpoint|

## 📦 Usage libraries & languages:

- [Go](https://go.dev/)

- [TypeScript](https://www.typescriptlang.org/)

|Language|Package|
|:---:|:---:|
|Go|[Gin](https://gin-gonic.com/)|
|GO|[Gin/cors](github.com/gin-contrib/cors)|
|Go|[Gin-Swagger](https://github.com/swaggo/gin-swagger)|
|TypeScript|[NextJS](https://nextjs.org/)|
|TypeScript|[tailwindcss](https://tailwindcss.com/)|
|TypeScript|[tabler/icons-react](https://tabler.io/icons)|