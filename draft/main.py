import requests
from bs4 import BeautifulSoup

# Title
# Company
# Location
# Link Company
# Description
# Email
# Salary
job = "golang"
location = "hanoi"
page = 1
url = f"https://www.topcv.vn/tim-viec-lam-{job}?sba={page}"
url1 = "https://www.linkedin.com/jobs/search/?currentJobId=4007749799&distance=25&geoId=104195383&keywords=Go%20(Programming%20Language)&origin=JOBS_HOME_KEYWORD_HISTORY&refresh=true"


def parser_html(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Accept-Language": "en-US,en;q=0.9",
    }
    response = requests.get(url, headers=headers)
    return BeautifulSoup(response.text, "html.parser")

def extract_link_from_a_tag(soup):
    content = ""
    if soup:
        if soup.name == "a":
            content = soup.get("href")
        else:
            link = soup.find("a")
            if link:
                content = link.get("href")
    return content

def extract_data(url, configs):
    soup = parser_html(url)
    jobs = soup.find_all(
        configs.get("find_jobs_tag"), class_=configs.get("find_jobs_class")
    )
    data = []
    if len(jobs) == 0:
        return data
    for job in jobs:
        title = job.find(configs.get("title_tag"), class_=configs.get("title_class"))
        company_name = job.find(
            configs.get("company_tag"), class_=configs.get("company_class")
        )
        location = job.find(
            configs.get("location_tag"), class_=configs.get("location_class")
        )
        salary = job.find(configs.get("salary_tag"), class_=configs.get("salary_class"))
        company_url = extract_link_from_a_tag(company_name)
        job_description_url = extract_link_from_a_tag(title)
        data.append(
            {
                "Title": title.text.strip(),
                "Company": company_name.text.strip(),
                "Company_URL": company_url,
                "Location": location.text.strip(),
                "Job Description": job_description_url,
                "Salary": salary.text.strip() if salary else "",
            }
        )
    return data


def main():
    topcv = {
        "find_jobs_tag": "div",
        "find_jobs_class": "job-item-search-result",
        "title_tag": "h3",
        "title_class": "title",
        "company_tag": "a",
        "company_class": "company",
        "location_tag": "label",
        "location_class": "address",
        "salary_tag": "label",
        "salary_class": "title-salary",
    }

    linkedin = {
        "find_jobs_tag": "div",
        "find_jobs_class": "base-card relative w-full hover:no-underline focus:no-underline base-card--link base-search-card base-search-card--link job-search-card",
        "title_tag": "a",
        "title_class": "base-card__full-link absolute top-0 right-0 bottom-0 left-0 p-0 z-[2]",
        "company_tag": "h4",
        "company_class": "base-search-card__subtitle",
        "location_tag": "span",
        "location_class": "job-search-card__location",
        "salary_tag": "div",
        "salary_class": "salary",
    }
    data = extract_data(url1, linkedin)
    print(data)


if __name__ == "__main__":
    main()
