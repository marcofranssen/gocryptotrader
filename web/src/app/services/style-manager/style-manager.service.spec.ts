import {inject, TestBed} from '@angular/core/testing';
import {HttpModule} from '@angular/http';
import {StyleManagerComponent} from './style-manager.component';
 

describe('StyleManager', () => {
  let styleManager: StyleManagerComponent;

  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpModule],
    providers: [StyleManagerComponent]
  }));

  beforeEach(inject([StyleManagerComponent], (sm: StyleManagerComponent) => {
    styleManager = sm;
  }));

  afterEach(() => {
    let links = document.head.querySelectorAll('link');
    for (let link of Array.prototype.slice.call(links)) {
      if (link.className.includes('style-manager-')) {
        document.head.removeChild(link);
      }
    }
  });

  it('should add stylesheet to head', () => {
    styleManager.setStyle('test', 'test.css');
    let styleEl = document.head.querySelector('.style-manager-test') as HTMLLinkElement;
    expect(styleEl).not.toBeNull();
    expect(styleEl.href.endsWith('test.css')).toBe(true);
  });

  it('should change existing stylesheet', () => {
    styleManager.setStyle('test', 'test.css');
    let styleEl = document.head.querySelector('.style-manager-test') as HTMLLinkElement;
    expect(styleEl).not.toBeNull();
    expect(styleEl.href.endsWith('test.css')).toBe(true);

    styleManager.setStyle('test', 'new.css');
    expect(styleEl.href.endsWith('new.css')).toBe(true);
  });

  it('should remove existing stylesheet', () => {
    styleManager.setStyle('test', 'test.css');
    let styleEl = document.head.querySelector('.style-manager-test') as HTMLLinkElement;
    expect(styleEl).not.toBeNull();
    expect(styleEl.href.endsWith('test.css')).toBe(true);

    styleManager.removeStyle('test');
    styleEl = document.head.querySelector('.style-manager-test') as HTMLLinkElement;
    expect(styleEl).toBeNull();
  });
});