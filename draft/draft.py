import requests
from bs4 import BeautifulSoup

# LinkedIn URL (example, use a real one if you have permission)
url = "https://www.linkedin.com/jobs/search/?currentJobId=4007749799&distance=25&geoId=104195383&keywords=Go%20(Programming%20Language)&origin=JOBS_HOME_KEYWORD_HISTORY&refresh=true"

# Custom headers (simulate a browser request)
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
    "Accept-Language": "en-US,en;q=0.9",
}

# Fetch the page
response = requests.get(url, headers=headers)

# Parse the HTML content using BeautifulSoup
soup = BeautifulSoup(response.text, 'html.parser')
a = soup.find_all('div', class_="base-card relative w-full hover:no-underline focus:no-underline base-card--link base-search-card base-search-card--link job-search-card")
#print(soup.prettify())

for job in a:
    #print(job)
    #title =  job.find('a', class_="base-card__full-link absolute top-0 right-0 bottom-0 left-0 p-0 z-[2]").text.strip()
    #location = job.find('span', class_="job-search-card__location").text.strip()
    #company = job.find('h4', class_="base-search-card__subtitle").text.strip()
    print(company)
# # Example: Extract the profile headline
# headline = soup.find("h1", {"class": "top-card-layout__title"}).text.strip()
#
# # Example: Extract work experience
# work_experience = soup.find_all("span", {"class": "mr1 t-bold"})
# for experience in work_experience:
#     print(experience.get_text().strip())

