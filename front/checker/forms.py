from django import forms

class UrlForm(forms.Form):
    url = forms.URLField(label='Enter a URL', widget=forms.URLInput(attrs={'placeholder': 'https://example.com'}))
