import requests
from django.shortcuts import render
from .forms import UrlForm

def check_url(request):
    result = None
    response_time = None

    if request.method == 'POST':
        form = UrlForm(request.POST)
        if form.is_valid():
            url = form.cleaned_data['url']
            try:
                # ارسال درخواست به بک‌اند Golang
                backend_response = requests.post(
                    "http://health-checker-service:8080/api/check",  # آدرس سرویس Golang در کوبرنتیس
                    json={"url": url},
                    timeout=5
                )
                data = backend_response.json()

                if backend_response.status_code == 200:
                    status_code = data.get("status_code")
                    response_time = data.get("response_time_ms")
                    result = f"✅ وضعیت سالم - کد پاسخ: {status_code}"
                else:
                    result = f"⚠️ خطا از سمت بک‌اند: {data.get('error', 'بدون توضیح')}"
            except requests.exceptions.RequestException as e:
                result = f"❌ خطا در ارتباط با بک‌اند: {str(e)}"
    else:
        form = UrlForm()

    return render(request, 'checker/form.html', {
        'form': form,
        'result': result,
        'response_time': response_time
    })
