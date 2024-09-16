import requests
from bs4 import BeautifulSoup
import re
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
pattern = r'(\d+)\+?\s*years?\s*of\s*experience'
# Parse the HTML content using BeautifulSoup
soup = BeautifulSoup(response.text, 'html.parser')
# a = soup.find('div', class_="job-detail__information-detail--content")
# print(a.prettify())
# if a:
#     all_text = a.get_text(separator='\n', strip=True)
#     print(all_text.split("Yêu cầu ứng viên")[-1])


fillter = {
    "min_exp": 3,
    "skill": ["python", "go", "nodejs"]
}


def regex(text, pattern):
    matches = re.findall(pattern, text, re.IGNORECASE)
    if matches:
        return matches.sort()
    return []


def rcm(text, fillter):
   skills = fillter.get("skill", [])
   min_exp = fillter.get("min_exp", 0)

   min_e = regex(text, pattern)
   if min_e and len(min_e) > 0 and min_e[-1] >= min_exp:
       return True
   if any(skill.lower() in text.lower() for skill in skills):
       return True
    
   return False

a = soup.find('div', class_="show-more-less-html__markup show-more-less-html__markup--clamp-after-5 relative overflow-hidden")
if a:
    all_text = a.get_text(separator='\n', strip=True).split("Requirements")
    if len(all_text) > 0:
        count = 0
        for text in all_text[-1].split("\n"):
            a = rcm(text, fillter)
            if a:
                count += 1
        if count >= len(fillter):
            print(True) 
        else:
            print(False)

# # Example: Extract the profile headline
# headline = soup.find("h1", {"class": "top-card-layout__title"}).text.strip()
#
# # Example: Extract work experience
# work_experience = soup.find_all("span", {"class": "mr1 t-bold"})
# for experience in work_experience:
#     print(experience.get_text().strip())
