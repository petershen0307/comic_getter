import requests
from bs4 import BeautifulSoup
import os

COMIC_SOURCE_URL = 'http://comic.sfacg.com'


def get_comic_by_volume(comic_value, comic_volume):
    pic_source_file = get_pic_array(comic_value, comic_volume)
    pic_paths = parse_pic_path(pic_source_file)
    target_folder = os.path.join(comic_value, comic_volume)
    if not os.path.exists(comic_value):
        os.mkdir(comic_value)
    if not os.path.exists(target_folder):
        os.mkdir(target_folder)
        download_pics(target_folder, pic_paths)


def download_pics(output_path, pic_paths):
    for index, url in enumerate(pic_paths, start=1):
        print('get url:{url}'.format(url=url))
        response = requests.get(url, stream=True)
        zero_fill = len(str(len(pic_paths)))
        file_name = str(index).zfill(zero_fill) + '.' + url.split('.')[-1]
        with open(os.path.join(output_path, file_name), 'wb') as f:
            for chunk in response.iter_content(chunk_size=1024):
                if chunk:
                    f.write(chunk)
        print('{file} Downloaded!'.format(file=file_name))


def get_pic_array(comic_value, comic_volume):
    assert isinstance(comic_value, str) or isinstance(comic_volume, str)
    http_request = requests.get(COMIC_SOURCE_URL + '/'.join(('', 'HTML', comic_value, comic_volume)))
    # print(http_request.url)
    # print(http_request.text)
    soup = BeautifulSoup(http_request.text, 'html.parser')
    for js_script in soup.find_all('script'):
        pic_source_file = js_script.get('src')
        # print(type(pic_source_file))
        # print(pic_source_file)
        if pic_source_file is not None and '.'.join((comic_volume, 'js')) in pic_source_file:
            print(pic_source_file)
            return pic_source_file


def parse_pic_path(file_name):
    http_request = requests.get(COMIC_SOURCE_URL + file_name)
    pic_path_collect = []
    for element in http_request.text.split(';'):
        pic_source = element.split('=')
        if 'picAy' in pic_source[0] and '/' in pic_source[1]:
            print(pic_source[1].replace('"', '').replace('\'', '').split()[0])
            pic_path_collect.append(COMIC_SOURCE_URL + pic_source[1].replace('"', '').replace('\'', '').split()[0])
    return pic_path_collect


if '__main__' == __name__:
    get_comic_by_volume('ASJS', '172')
