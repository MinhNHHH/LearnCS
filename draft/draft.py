import requests
from bs4 import BeautifulSoup

# LinkedIn URL (example, use a real one if you have permission)
url = "https://www.linkedin.com/jobs/view/golang-developer-can-work-remotely-at-sotatek-jsc-4007752495/"
url1 = "https://www.topcv.vn/viec-lam/senior-golang-developer/1376241.html?ta_source=JobSearchList_LinkDetail&u_sr_id=KiSczNgZGvWTmLlECuYzPOwKJDnt4F2NtxlhP0Xg_1726470930"
# Custom headers (simulate a browser request)
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
    "Accept-Language": "en-US,en;q=0.9",
}

# Fetch the page
response = requests.get(url, headers=headers)

# Parse the HTML content using BeautifulSoup
soup = BeautifulSoup(response.text, 'html.parser')
# a = soup.find('div', class_="job-detail__information-detail--content")
# print(a.prettify())
# if a:
#     all_text = a.get_text(separator='\n', strip=True)
#     print(all_text.split("Yêu cầu ứng viên")[-1])


fillter = {
    "min-exp": 3,
    "skill": ["python", "go", "nodejs"]
}


# def fillter(text, condition):
#
#
# a = soup.find('div', class_="show-more-less-html__markup show-more-less-html__markup--clamp-after-5 relative overflow-hidden")
# if a:
#     all_text = a.get_text(separator='\n', strip=True).split("Requirements")
#     if len(all_text) > 0:
#         for text in all_text[-1].split("\n"):
#

import re
 # Sample text containing experience information
text = "The candidate has 3+ years of experience with Go development and 5 years with Python."

# Regular expression pattern to match experience information
pattern = r'(\d+)\+?\s*years?\s*of\s*experience'

# Find all matches
matches = re.findall(pattern, text, re.IGNORECASE)   
print(matches)

# # Example: Extract the profile headline
# headline = soup.find("h1", {"class": "top-card-layout__title"}).text.strip()
#
# # Example: Extract work experience
# work_experience = soup.find_all("span", {"class": "mr1 t-bold"})
# for experience in work_experience:
#     print(experience.get_text().strip())
