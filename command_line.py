import configparser
import getter


def start():
    ini_parser = configparser.ConfigParser()
    ini_parser.read('setting.ini', encoding='utf8')
    # get url
    url = 'http://comic.sfacg.com/HTML/iam/'
    # get url
    if 'HTML' not in url:
        print('not support url:({}) right now'.format(url))
    if url[-1] is '/':
        url = url[:-1]
    getter.get_comic_by_volume(url.split('/')[-1], '001')

if '__main__' == __name__:
    start()
