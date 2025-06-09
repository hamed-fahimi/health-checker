# Health Checker

Health Checker یک ابزار مانیتورینگ ساده برای بررسی وضعیت سرویس‌های HTTP/HTTPS است که به صورت دوره‌ای آدرس‌های مشخص‌شده را چک می‌کند و نتایج را ذخیره می‌کند.

---

## ویژگی‌ها

- بررسی سلامت سرویس‌ها با درخواست HTTP GET به صورت زمان‌بندی‌شده
- ذخیره نتایج بررسی‌ها در پایگاه داده SQLite
- API ساده (با Gin) برای مشاهده وضعیت سرویس‌ها (`/status`)
- توسعه‌پذیر برای افزودن سرویس‌های جدید
- قابلیت اجرا به صورت محلی و در Docker
- **پشتیبانی از استقرار روی Kubernetes**
- **دارای رابط کاربری فرانت‌اند با Django (پوشه `front/`)**
- **Dockerfile برای بک‌اند و فرانت‌اند جهت داکرایز کردن کامل پروژه**

---

## ساختار پروژه

```text
health-checker/
├── cmd/                  # نقطه شروع برنامه Go
├── internal/
│   ├── api/              # API و Router
│   ├── checker/          # منطق Health Check
│   └── db/               # مدل و اتصال دیتابیس
├── Dockerfile            # Dockerfile بک‌اند
├── go.mod                # وابستگی‌های Go
├── README.md
└── ...
```

---

## نحوه اجرا

### اجرای محلی

1. کلون و نصب وابستگی‌ها:
   ```bash
   git clone https://github.com/hamed-fahimi/health-checker.git
   cd health-checker
   go mod tidy
   ```

2. اجرای برنامه:
   ```bash
   go run cmd/main.go
   ```

3. تست API:
   ```bash
   curl http://localhost:8080/status
   ```

---

### اجرای با Docker

1. ساخت ایمیج:
   ```bash
   docker build -t health-checker .
   ```

2. اجرای کانتینر:
   ```bash
   docker run -p 8080:8080 health-checker
   ```

---

## استقرار روی Kubernetes

در صورت نیاز به استقرار روی کوبرنتیز، فایل‌های yaml مربوط به Deployment و Service را در پوشه `k8s/` قرار دهید و با دستور زیر اجرا کنید:

```bash
kubectl apply -f k8s/
```

---

## فرانت‌اند (Django)

پوشه `front/` شامل پروژه Django است که رابط کاربری برای مشاهده و مدیریت وضعیت سرویس‌ها را فراهم می‌کند.  
برای اجرای فرانت‌اند:

```bash
cd front
python -m venv venv
source venv/bin/activate  # یا در ویندوز: venv\Scripts\activate
pip install -r requirements.txt
python manage.py migrate
python manage.py runserver 0.0.0.0:8000
```

همچنین می‌توانید فرانت‌اند را با Docker اجرا کنید:

```bash
docker build -t health-checker-frontend -f dockerfile .
docker run -p 8000:8000 health-checker-frontend
```

---

## مشارکت

1. Fork کنید.
2. Branch جدید بسازید.
3. تغییرات را Commit و Push کنید.
4. Pull Request ارسال کنید.

---

## لایسنس

MIT

---

## ارتباط

[hamed-fahimi](https://github.com/hamed-fahimi)  
hamedfahimi3@gmail.com