import configparser
import requests
from bs4 import BeautifulSoup


# find comic name and url in <table></table>
def parse_comic_name(html_content, section=0):
    soup = BeautifulSoup(html_content, 'html.parser')
    # tag -> table, attribute -> comic_cover Height_px22 font_gray
    table = soup.find('table', 'comic_cover Height_px22 font_gray')
    td = table.find('td')
    comic_url_map = configparser.ConfigParser()
    for index, li in enumerate(td.find_all('li', 'Conjunction'), start=section):
        a = li.find('a')
        img = li.find('img')
        # get('attribute') -> get attribute value
        href = a.get('href')
        comic_name = img.get('alt')
        # print(index, ' url:', href, ' comic name:', comic_name)
        # .replace('%', '%%') is workaround for '%' at value [Interpolation in configparser]
        comic_url_map[index] = {'name': comic_name.replace('%', '%%'), 'url': href}
    with open('setting.ini', mode='a', encoding='utf8') as config_file:
        comic_url_map.write(config_file)
    return index + 1  # next start section


def parse_total_page(html_content):
    soup = BeautifulSoup(html_content, 'html.parser')
    li = soup.find('li', 'pagebarNext')
    if li is None:
        return None
    a = li.find('a')
    url = a.get('href')
    # print(url)
    return url


def query_url():
    start_key = 0
    url = 'http://comic.sfacg.com/Catalog/'
    while True:
        print('current url:', url)
        http_request = requests.get(url)
        html_content = http_request.text
        start_key = parse_comic_name(html_content, start_key)
        url = parse_total_page(html_content)
        if url is None:
            break

if '__main__' == __name__:
    query_url()
    # with open('test_name.html', mode='r', encoding='utf8') as f:
    #     # parse_comic_name(f.read())
    #     parse_total_page(f.read())
